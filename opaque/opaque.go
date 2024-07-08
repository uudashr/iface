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
		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return
		}

		if fn.Type.Results == nil {
			return
		}

		if fn.Recv != nil {
			return
		}

		if debug {
			fmt.Println("Function declaration:", fn.Name.Name, fn.Pos(), len(fn.Type.Results.List))
		}

		returnTypes := make([]types.Type, len(fn.Type.Results.List))
		hasIfaceReturnType := false

		for i, field := range fn.Type.Results.List {
			if debug {
				fmt.Printf(" [%d] Field: %v %v\n", i, field.Type, reflect.TypeOf(field.Type))
			}

			typ := pass.TypesInfo.TypeOf(field.Type)
			if typ == nil {
				continue
			}

			returnTypes[i] = typ

			_, ok := typ.Underlying().(*types.Interface)
			if !ok {
				continue
			}

			if debug {
				fmt.Printf("     Return type is an interface, samePackage: %t\n", fromSamePackage(pass, typ))
			}

			hasIfaceReturnType = true
		}

		if debug {
			fmt.Println(" Return types:", returnTypes, hasIfaceReturnType)
		}

		if !hasIfaceReturnType {
			return
		}

		concreteStmtTypes := make([]map[types.Type]struct{}, len(fn.Type.Results.List))
		for i := range concreteStmtTypes {
			concreteStmtTypes[i] = make(map[types.Type]struct{})
		}

		// Inspect return statements within the function body
		ast.Inspect(fn.Body, func(n ast.Node) bool {
			switch n := n.(type) {
			case *ast.FuncLit:
				return false
			case *ast.ReturnStmt:
				if debug {
					fmt.Println(" Return statement:", n.Pos(), len(n.Results))
				}

				for i, exp := range n.Results {
					if i >= len(returnTypes) {
						continue
					}

					stmtTyp := pass.TypesInfo.TypeOf(exp)
					if stmtTyp == nil {
						continue
					}

					if debug {
						fmt.Printf("  [%d] %v %v %v\n", i, reflect.TypeOf(exp), stmtTyp.Underlying(), reflect.TypeOf(stmtTyp.Underlying()))
					}

					_, ifaceRetType := returnTypes[i].Underlying().(*types.Interface)
					if !ifaceRetType {
						continue
					}

					_, ok = stmtTyp.Underlying().(*types.Interface)
					if ok {
						continue
					}

					basic, ok := stmtTyp.(*types.Basic)
					if ok && basic.Kind() == types.UntypedNil {
						// ignore nil return statement
						continue
					}

					_, ok = stmtTyp.Underlying().(*types.Signature)
					if ok {
						// ignore function type return statement
						continue
					}

					concreteStmtTypes[i][stmtTyp] = struct{}{}

					if debug {
						fmt.Printf("   Return statement concrete %v, return type is interface\n", stmtTyp)
					}
				}
			}

			return true
		})

		for i, retType := range returnTypes {
			_, ifaceRetType := retType.Underlying().(*types.Interface)
			if !ifaceRetType {
				continue
			}

			if !fromSamePackage(pass, retType) {
				continue
			}

			concretes := len(concreteStmtTypes[i])
			if concretes == 0 {
				// no concrete statement return
				continue
			}

			if concretes > 1 {
				// has multiple concrete statement returns
				continue
			}

			var concrete types.Type
			for c := range concreteStmtTypes[i] {
				concrete = c
			}

			// has a single concrete statement return
			pos := fn.Type.Results.List[i].Pos()

			pass.Reportf(pos, "%s function return %s at the %s result, abstract a single concrete implementation of %s",
				fn.Name.Name, removePkgPrefix(retType.String()), toHumanOrdinal(i), concrete)
		}
	})

	return nil, nil
}

func fromSamePackage(pass *analysis.Pass, typ types.Type) bool {
	currentPkg := pass.Pkg

	named, ok := typ.(*types.Named)
	if !ok {
		return false
	}

	ifacePkg := named.Obj().Pkg()

	return currentPkg == ifacePkg
}

func removePkgPrefix(typeStr string) string {
	if lastDot := strings.LastIndex(typeStr, "."); lastDot != -1 {
		return typeStr[lastDot+1:]
	}

	return typeStr
}

func toHumanOrdinal(index int) string {
	switch index {
	case 0:
		return "1st"
	case 1:
		return "2nd"
	case 2:
		return "3rd"
	}

	return fmt.Sprintf("%dth", index+1)
}
