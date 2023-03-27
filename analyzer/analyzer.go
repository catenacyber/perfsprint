package analyzer

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:     "gostrconv",
	Doc:      "Checks that fmt.Sprintf can be replaced with a faster strconv function.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	var fmtpkg *types.Package
	for _, pkg := range pass.Pkg.Imports() {
		if pkg.Path() == "fmt" {
			fmtpkg = pkg
		}
	}
	if fmtpkg == nil {
		return nil, nil
	}
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
		if pass.TypesInfo.ObjectOf(called.Sel).Pkg() != fmtpkg {
			return
		}
		if called.Sel.Name != "Sprintf" {
			return
		}
		if len(call.Args) != 2 {
			return
		}
		arg0, ok := call.Args[0].(*ast.BasicLit)
		if !ok {
			return
		}
		var base int
		switch arg0.Value {
		case `"%d"`:
			base = 10
		case `"%x"`:
			base = 16
		default:
			return
		}
		v := pass.TypesInfo.TypeOf(call.Args[1])
		if types.Identical(v, types.Typ[types.Int]) && base == 10 {
			pass.Report(analysis.Diagnostic{
				Pos:     node.Pos(),
				End:     node.End(),
				Message: "fmt.Sprintf can be replaced with faster function strconv.Itoa",
				// need ro run goimports to fix use of fmt/strconv afterwards
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: "use strconv.Itoa",
						TextEdits: []analysis.TextEdit{{
							Pos:     node.Pos(),
							End:     call.Args[1].Pos(),
							NewText: []byte("strconv.Itoa("),
						}},
					},
				},
			})
		} else if types.Identical(v, types.Typ[types.Int64]) {
			pass.Report(analysis.Diagnostic{
				Pos:     node.Pos(),
				End:     node.End(),
				Message: "fmt.Sprintf can be replaced with faster function strconv.FormatInt",
				// need ro run goimports to fix use of fmt/strconv afterwards
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: "use strconv.FormatInt",
						TextEdits: []analysis.TextEdit{{
							Pos:     node.Pos(),
							End:     call.Args[1].Pos(),
							NewText: []byte("strconv.FormatInt("),
						},
							{
								Pos:     call.Args[1].End(),
								End:     call.Args[1].End(),
								NewText: []byte(fmt.Sprintf(", %d", base)),
							},
						},
					},
				},
			})
		} else if types.Identical(v, types.Typ[types.Uint64]) {
			pass.Report(analysis.Diagnostic{
				Pos:     node.Pos(),
				End:     node.End(),
				Message: "fmt.Sprintf can be replaced with faster function strconv.FormatUint",
				// need ro run goimports to fix use of fmt/strconv afterwards
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: "use strconv.FormatUint",
						TextEdits: []analysis.TextEdit{{
							Pos:     node.Pos(),
							End:     call.Args[1].Pos(),
							NewText: []byte("strconv.FormatUint("),
						},
							{
								Pos:     call.Args[1].End(),
								End:     call.Args[1].End(),
								NewText: []byte(fmt.Sprintf(", %d", base)),
							},
						},
					},
				},
			})
		} else {
			pass.Reportf(node.Pos(), "Sprintf can be replaced with faster function from strconv")
		}
	})

	return nil, nil
}
