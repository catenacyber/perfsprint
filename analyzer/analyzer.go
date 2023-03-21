package analyzer

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:     "gostrconv",
	Doc:      "Checks that Sprintf can be replaced with a faster strconv function.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		call := node.(*ast.CallExpr)
		called, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return
		}
		if fmt.Sprintf("%s.%s", called.X, called.Sel) != "fmt.Sprintf" {
			return
		}
		if len(call.Args) != 2 {
			return
		}
		arg0, ok := call.Args[0].(*ast.BasicLit)
		if !ok {
			return
		}
		if arg0.Value != `"%d"` {
			return
		}

		pass.Reportf(node.Pos(), "Sprintf can be replaced with faster function from strconv")
	})

	return nil, nil
}
