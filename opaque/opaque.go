package opaque

import (
	"fmt"
	"go/ast"
	"go/types"
	"os"
	"strings"

	"github.com/uudashr/iface/internal/directive"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer is the analysis pass for detecting opaque interface returns.
var Analyzer = newAnalyzer()

func newAnalyzer() *analysis.Analyzer {
	r := runner{}

	analyzer := &analysis.Analyzer{
		Name:     "opaque",
		Doc:      "Detects functions that return an interface type, but only ever return a single concrete implementation.",
		URL:      "https://pkg.go.dev/github.com/uudashr/iface/opaque",
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run:      r.run,
	}

	analyzer.Flags.BoolVar(&r.debug, "nerd", false, "enable nerd mode")

	return analyzer
}

type runner struct {
	debug bool
}

func (r *runner) run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// Find function declarations that return an interface

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		funcDecl := n.(*ast.FuncDecl)

		if funcDecl.Recv != nil {
			// skip methods
			return
		}

		if funcDecl.Body == nil {
			// skip functions without body
			return
		}

		if funcDecl.Type.Results == nil {
			// skip functions without return values
			return
		}

		if r.debug {
			fmt.Fprintf(os.Stderr, "Function declaration %s\n", funcDecl.Name.Name)
			fmt.Fprintf(os.Stderr, " Results len=%d\n", len(funcDecl.Type.Results.List))
		}

		dir := directive.ParseIgnore(funcDecl.Doc)
		if dir != nil && dir.ShouldIgnore(pass.Analyzer.Name) {
			// skip ignored function
			return
		}

		// Pre-check, only function that has interface return type will be processed
		var hasInterfaceReturnType bool

		var outCount int

		for i, result := range funcDecl.Type.Results.List {
			outInc := 1
			if namesLen := len(result.Names); namesLen > 0 {
				outInc = namesLen
			}

			outCount += outInc

			resType := result.Type
			typ := pass.TypesInfo.TypeOf(resType)

			if r.debug {
				fmt.Fprintf(os.Stderr, "  [%d] len=%d %v %v %T | %v %T interface=%t\n", i, len(result.Names), result.Names, resType, resType, typ, typ, types.IsInterface(typ))
			}

			if types.IsInterface(typ) && !hasInterfaceReturnType {
				hasInterfaceReturnType = true
			}
		}

		r.debugf("  hasInterface=%t outCount=%d\n", hasInterfaceReturnType, outCount)

		if !hasInterfaceReturnType {
			// skip, since it has no interface return type
			return
		}

		// Collect types on every return statement
		retStmtTypes := make([]map[types.Type]struct{}, outCount)
		for i := range retStmtTypes {
			retStmtTypes[i] = make(map[types.Type]struct{})
		}

		ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
			switch n := n.(type) {
			case *ast.FuncLit:
				// ignore nested functions
				return false
			case *ast.ReturnStmt:
				if r.debug {
					fmt.Fprintf(os.Stderr, "  Return statements %v len=%d\n", n.Results, len(n.Results))
				}

				for i, result := range n.Results {
					r.debugf("   [%d] %v %T\n", i, result, result)

					switch res := result.(type) {
					case *ast.CallExpr:
						r.debugf("       CallExpr Fun: %v %T\n", res.Fun, res.Fun)

						typ := pass.TypesInfo.TypeOf(res)
						switch typ := typ.(type) {
						case *types.Tuple:
							for j := range typ.Len() {
								v := typ.At(j)
								vTyp := v.Type()
								retStmtTypes[j][vTyp] = struct{}{}

								if r.debug {
									fmt.Fprintf(os.Stderr, "          Tuple [%d]: %v %T | %v %T interface=%t\n", j, v, v, vTyp, vTyp, types.IsInterface(vTyp))
								}
							}
						default:
							retStmtTypes[i][typ] = struct{}{}
						}

					case *ast.Ident:
						r.debugf("       Ident: %v %T\n", res, res)

						typ := pass.TypesInfo.TypeOf(res)
						isNilStmt := isUntypedNil(typ)

						if r.debug {
							fmt.Fprintf(os.Stderr, "        Ident type: %v %T interface=%t, untypedNil=%t\n", typ, typ, types.IsInterface(typ), isNilStmt)
						}

						if !isNilStmt {
							retStmtTypes[i][typ] = struct{}{}
						}
					case *ast.UnaryExpr:
						r.debugf("       UnaryExpr X: %v \n", res.X)

						typ := pass.TypesInfo.TypeOf(res)

						if r.debug {
							fmt.Fprintf(os.Stderr, "        UnaryExpr type: %v %T interface=%t\n", typ, typ, types.IsInterface(typ))
						}

						retStmtTypes[i][typ] = struct{}{}
					default:
						r.debugf("       Unknown: %v %T\n", res, res)

						typ := pass.TypesInfo.TypeOf(res)
						retStmtTypes[i][typ] = struct{}{}
					}
				}

				return false
			default:
				return true
			}
		})

		// Compare func return types with the return statement types
		var nextIdx int

		for _, result := range funcDecl.Type.Results.List {
			resType := result.Type
			typ := pass.TypesInfo.TypeOf(resType)

			consumeCount := 1
			if len(result.Names) > 0 {
				consumeCount = len(result.Names)
			}

			currentIdx := nextIdx
			nextIdx += consumeCount

			// Check return type
			if !types.IsInterface(typ) {
				// it is a concrete type
				continue
			}

			errorType := types.Universe.Lookup("error").Type()
			anyType := types.Universe.Lookup("any").Type()

			if types.Identical(typ, errorType) || types.Identical(typ, anyType) {
				continue
			}

			if !fromSamePackage(pass, typ) {
				// ignore interface from other package
				continue
			}

			// Check statement type
			stmtTyps := retStmtTypes[currentIdx]

			stmtTypsSize := len(stmtTyps)
			if stmtTypsSize > 1 {
				// it has multiple implementation
				continue
			}

			if stmtTypsSize == 0 {
				// function use named return value, while return statement is empty
				continue
			}

			var stmtTyp types.Type
			for t := range stmtTyps {
				// expect only one, we don't have to break it
				stmtTyp = t
			}

			if types.IsInterface(stmtTyp) {
				// not concrete type, skip
				continue
			}

			if r.debug {
				fmt.Fprintf(os.Stderr, "stmtType: %v %T | %v %T\n", stmtTyp, stmtTyp, stmtTyp.Underlying(), stmtTyp.Underlying())
			}

			switch stmtTyp := stmtTyp.(type) {
			case *types.Basic:
				if stmtTyp.Kind() == types.UntypedNil {
					// ignore nil
					continue
				}
			case *types.Named:
				if _, ok := stmtTyp.Underlying().(*types.Signature); ok {
					// skip function type
					continue
				}
			}

			retTypeName := typ.String()
			if fromSamePackage(pass, typ) {
				retTypeName = removePkgPrefix(retTypeName)
			}

			stmtTypName := stmtTyp.String()
			if fromSamePackage(pass, stmtTyp) {
				stmtTypName = removePkgPrefix(stmtTypName)
			}

			msg := fmt.Sprintf("'%s' function return '%s' interface at the %s result, abstract a single concrete implementation of '%s'",
				funcDecl.Name.Name,
				retTypeName,
				positionStr(currentIdx),
				stmtTypName)

			pass.Report(analysis.Diagnostic{
				Pos:     result.Pos(),
				Message: msg,
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: "Replace the interface return type with the concrete type",
						TextEdits: []analysis.TextEdit{
							{
								Pos:     result.Pos(),
								End:     result.End(),
								NewText: []byte(stmtTypName),
							},
						},
					},
				},
			})
		}
	})

	return nil, nil
}

func isUntypedNil(typ types.Type) bool {
	if b, ok := typ.(*types.Basic); ok {
		return b.Kind() == types.UntypedNil
	}

	return false
}

func positionStr(idx int) string {
	switch idx {
	case 0:
		return "1st"
	case 1:
		return "2nd"
	case 2:
		return "3rd"
	default:
		return fmt.Sprintf("%dth", idx+1)
	}
}

func fromSamePackage(pass *analysis.Pass, typ types.Type) bool {
	switch typ := typ.(type) {
	case *types.Named:
		currentPkg := pass.Pkg
		ifacePkg := typ.Obj().Pkg()

		return currentPkg == ifacePkg
	case *types.Pointer:
		return fromSamePackage(pass, typ.Elem())
	default:
		return false
	}
}

func removePkgPrefix(typeStr string) string {
	if typeStr[0] == '*' {
		return "*" + removePkgPrefix(typeStr[1:])
	}

	if lastDot := strings.LastIndex(typeStr, "."); lastDot != -1 {
		return typeStr[lastDot+1:]
	}

	return typeStr
}

func (r *runner) debugf(format string, a ...any) {
	if r.debug {
		fmt.Fprintf(os.Stderr, format, a...)
	}
}
