package opaque

import (
	"fmt"
	"go/ast"
	"go/types"
	"reflect"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var debug = false

func init() {
	Analyzer.Flags.BoolVar(&debug, "debug", false, "enable debug mode")
}

var Analyzer = &analysis.Analyzer{
	Name:     "opaque",
	Doc:      "finds the interfaces that is used to abstract a single concrete implementation only",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
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

		if funcDecl.Type.Results == nil {
			// skip functions without return values
			return
		}

		if debug {
			fmt.Printf("Function declaration %s\n", funcDecl.Name.Name)
		}

		retLen := len(funcDecl.Type.Results.List)

		if debug {
			fmt.Printf(" Results %d\n", retLen)
		}

		// Pre-check, only function that has interface return type will be processed
		var hasInterfaceReturnType bool

		for i, result := range funcDecl.Type.Results.List {
			resType := result.Type
			typ := pass.TypesInfo.TypeOf(resType)

			if debug {
				fmt.Printf("  [%d] %v %v | %v %v interface=%t\n", i, resType, reflect.TypeOf(resType), typ, reflect.TypeOf(typ), types.IsInterface(typ))
			}

			if types.IsInterface(typ) && !hasInterfaceReturnType {
				hasInterfaceReturnType = true
			}
		}

		if !hasInterfaceReturnType {
			// skip, since it has no interface return type
			return
		}

		// Collect types on every return statement
		retStmtTypes := make([]map[types.Type]struct{}, retLen)
		for i := range retLen {
			retStmtTypes[i] = make(map[types.Type]struct{})
		}

		ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
			// fmt.Printf("  node: %v %v\n", n, reflect.TypeOf(n))
			switch n := n.(type) {
			case *ast.FuncLit:
				// ignore nested functions
				return false
			case *ast.ReturnStmt:
				if debug {
					fmt.Printf("  Return statements %v len=%d\n", n.Results, len(n.Results))
				}

				for i, result := range n.Results {
					if debug {
						fmt.Printf("   [%d] %v %v\n", i, result, reflect.TypeOf(result))
					}

					switch res := result.(type) {
					case *ast.CallExpr:
						if debug {
							fmt.Printf("       CallExpr Fun: %v %v\n", res.Fun, reflect.TypeOf(res.Fun))
						}

						typ := pass.TypesInfo.TypeOf(res)
						switch typ := typ.(type) {
						case *types.Tuple:
							for i := range typ.Len() {
								v := typ.At(i)
								vTyp := v.Type()
								retStmtTypes[i][vTyp] = struct{}{}

								if debug {
									fmt.Printf("          Tuple [%d]: %v %v | %v %v interface=%t\n", i, v, reflect.TypeOf(v), vTyp, reflect.TypeOf(vTyp), types.IsInterface(vTyp))
								}
							}
						default:
							retStmtTypes[i][typ] = struct{}{}
						}

					case *ast.Ident:
						if debug {
							fmt.Printf("       Ident: %v %v\n", res, reflect.TypeOf(res))
						}

						typ := pass.TypesInfo.TypeOf(res)

						if debug {
							fmt.Printf("        Ident type: %v %v interface=%t\n", typ, reflect.TypeOf(typ), types.IsInterface(typ))
						}

						retStmtTypes[i][typ] = struct{}{}
					case *ast.UnaryExpr:
						if debug {
							fmt.Printf("       UnaryExpr X: %v \n", res.X)
						}

						typ := pass.TypesInfo.TypeOf(res)

						if debug {
							fmt.Printf("        UnaryExpr type: %v %v interface=%t\n", typ, reflect.TypeOf(typ), types.IsInterface(typ))
						}

						retStmtTypes[i][typ] = struct{}{}
					default:
						if debug {
							fmt.Printf("       Unknown: %v %v\n", res, reflect.TypeOf(res))
						}

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
		for i, result := range funcDecl.Type.Results.List {
			resType := result.Type
			typ := pass.TypesInfo.TypeOf(resType)

			// Check return type
			if !types.IsInterface(typ) {
				// it is a concrete type
				continue
			}

			if typ.String() == "error" {
				// very common case to have return type error
				continue
			}

			if !fromSamePackage(pass, typ) {
				// ignore interface from other package
				continue
			}

			// Check statement type
			stmtTyps := retStmtTypes[i]

			stmtTypsSize := len(stmtTyps)
			if stmtTypsSize > 1 {
				// it has multiple implementation
				continue
			}

			if stmtTypsSize != 1 {
				panic("expect stmtTypsSize equal to 1")
			}

			var stmtTyp types.Type
			for t := range stmtTyps {
				stmtTyp = t
				// expect only one, we don't have to break it
			}

			if types.IsInterface(stmtTyp) {
				// not concrete type, skip
				continue
			}

			if debug {
				fmt.Printf("stmtType: %v %v | %v %v\n", stmtTyp, reflect.TypeOf(stmtTyp), stmtTyp.Underlying(), reflect.TypeOf(stmtTyp.Underlying()))
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

			pass.Reportf(result.Pos(),
				"%s function return %s interface at the %s result, abstract a single concrete implementation of %s",
				funcDecl.Name.Name,
				retTypeName,
				positionStr(i),
				stmtTypName)
		}
	})

	return nil, nil
}

func positionStr(i int) string {
	switch i {
	case 0:
		return "1st"
	case 1:
		return "2nd"
	case 2:
		return "3rd"
	default:
		return fmt.Sprintf("%dth", i+1)
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
