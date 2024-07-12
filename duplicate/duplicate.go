package duplicate

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"reflect"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var debug = false

func init() {
	Analyzer.Flags.BoolVar(&debug, "debug", false, "enable debug mode")
}

var Analyzer = &analysis.Analyzer{
	Name:     "duplicate",
	Doc:      "finds duplicate interfaces within the package",
	URL:      "https://pkg.go.dev/github.com/uudashr/iface/duplicate",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// Collect interface type declarations
	ifaceDecls := make(map[string]token.Pos)
	ifaceTypes := make(map[string]*types.Interface)

	nodeFilter := []ast.Node{
		(*ast.TypeSpec)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return
		}

		ifaceType, ok := ts.Type.(*ast.InterfaceType)
		if !ok {
			return
		}

		if debug {
			fmt.Println("Interface declaration:", ts.Name.Name, ts.Pos(), len(ifaceType.Methods.List))

			for i, field := range ifaceType.Methods.List {
				switch ft := field.Type.(type) {
				case *ast.FuncType:
					fmt.Printf(" [%d] Field: func %s %v %v\n", i, field.Names[0].Name, reflect.TypeOf(field.Type), field.Pos())
				case *ast.Ident:
					fmt.Printf(" [%d] Field: iface %s %v %v\n", i, ft.Name, reflect.TypeOf(field.Type), field.Pos())
				default:
					fmt.Printf(" [%d] Field: unknown %v\n", i, reflect.TypeOf(ft))
				}
			}
		}

		ifaceDecls[ts.Name.Name] = ts.Pos()

		obj := pass.TypesInfo.Defs[ts.Name]
		if obj == nil {
			return
		}

		iface, ok := obj.Type().Underlying().(*types.Interface)
		if !ok {
			return
		}

		ifaceTypes[ts.Name.Name] = iface
	})

Loop:
	for name, typ := range ifaceTypes {
		for otherName, otherTyp := range ifaceTypes {
			if name == otherName {
				continue
			}

			if !types.Identical(typ, otherTyp) {
				continue
			}

			if debug {
				fmt.Println("Duplicate interface:", name, "and", otherName)
			}

			pass.Reportf(ifaceDecls[name], "interface %s contains duplicate methods or type constraints from another interface, causing redundancy", name)

			continue Loop
		}
	}

	return nil, nil
}
