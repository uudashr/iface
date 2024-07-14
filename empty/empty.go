package empty

import (
	"fmt"
	"go/ast"
	"reflect"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer is the empty interface analyzer.
var Analyzer = newAnalyzer()

func newAnalyzer() *analysis.Analyzer {
	r := runner{}

	analyzer := &analysis.Analyzer{
		Name:     "empty",
		Doc:      "finds empty interfaces",
		URL:      "https://pkg.go.dev/github.com/uudashr/iface/empty",
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
			fmt.Println("Interface type declaration:", ts.Name.Name, ts.Pos(), len(ifaceType.Methods.List))

			for i, f := range ifaceType.Methods.List {
				fmt.Printf(" [%d] %v %v\n", i, f.Names, reflect.TypeOf(f.Type))
			}
		}

		if len(ifaceType.Methods.List) == 0 {
			pass.Reportf(ts.Pos(), "interface %s is empty, providing no meaningful behavior", ts.Name.Name)
		}
	})

	return nil, nil
}
