package unused

import (
	"fmt"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var debug = false

func init() {
	Analyzer.Flags.BoolVar(&debug, "debug", false, "enable debug mode")
}

var Analyzer = &analysis.Analyzer{
	Name:     "unused",
	Doc:      "finds unused interfaces within the package",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// Collect all interface type declarations
	ifaceDecls := make(map[string]token.Pos)

	nodeFilter := []ast.Node{
		(*ast.TypeSpec)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return
		}

		_, ok = ts.Type.(*ast.InterfaceType)
		if !ok {
			return
		}

		if debug {
			fmt.Println("Interface type declaration:", ts.Name.Name, ts.Pos())
		}

		ifaceDecls[ts.Name.Name] = ts.Pos()
	})

	if debug {
		var ifaceNames []string
		for name := range ifaceDecls {
			ifaceNames = append(ifaceNames, name)
		}

		fmt.Println("Declared interfaces:", ifaceNames)
	}

	// Inspect whether the interface is used within the package
	nodeFilter = []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		ident, ok := n.(*ast.Ident)
		if !ok {
			return
		}

		pos := ifaceDecls[ident.Name]
		if pos == ident.Pos() {
			// The identifier is the interface type declaration
			return
		}

		delete(ifaceDecls, ident.Name)
	})

	for name, pos := range ifaceDecls {
		pass.Reportf(pos, "interface %s is declared but not used within the package", name)
	}

	return nil, nil
}
