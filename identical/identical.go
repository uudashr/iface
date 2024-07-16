package identical

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

// Analyzer is the duplicate interface analyzer.
var Analyzer = newAnalyzer()

func newAnalyzer() *analysis.Analyzer {
	r := runner{}

	analyzer := &analysis.Analyzer{
		Name:     "duplicate",
		Doc:      "Identifies interfaces in the same package that have identical method sets",
		URL:      "https://pkg.go.dev/github.com/uudashr/iface/duplicate",
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run:      r.run,
	}

	analyzer.Flags.BoolVar(&r.debug, "debug", false, "enable debug mode")

	return analyzer
}

type runner struct {
	debug bool
}

func (r *runner) run(pass *analysis.Pass) (interface{}, error) {
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

		if r.debug {
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

			if r.debug {
				fmt.Println("Identical interface:", name, "and", otherName)
			}

			pass.Reportf(ifaceDecls[name], "interface %s contains identical methods or type constraints from another interface, causing redundancy", name)

			continue Loop
		}
	}

	return nil, nil
}
