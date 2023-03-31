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
		case `"%d"`, `"%v"`:
			base = 10
		case `"%x"`:
			base = 16
		case `"%t"`:
			base = 1
		case `"%s"`:
			base = 0
		default:
			return
		}
		v := pass.TypesInfo.TypeOf(call.Args[1])
		s, isslice := v.(*types.Slice)
		a, isarray := v.(*types.Array)
		if types.Identical(v, types.Typ[types.Float32]) || types.Identical(v, types.Typ[types.Float64]) {
		} else if types.Identical(v, types.Typ[types.String]) && base == 0 {
			pass.Report(analysis.Diagnostic{
				Pos:     node.Pos(),
				End:     node.End(),
				Message: "fmt.Sprintf can be replaced with just using the string",
				// need ro run goimports to fix use of fmt/strconv afterwards
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: "use strconv.FormatBool",
						TextEdits: []analysis.TextEdit{{
							Pos:     node.Pos(),
							End:     call.Args[1].Pos(),
							NewText: []byte(""),
						},
							{
								Pos:     call.Args[1].End(),
								End:     node.End(),
								NewText: []byte(""),
							},
						},
					},
				},
			})
		} else if base == 0 && v.String() == "error" {
			pass.Report(analysis.Diagnostic{
				Pos:     node.Pos(),
				End:     node.End(),
				Message: "fmt.Sprintf can be replaced with using Error()",
				// need ro run goimports to fix use of fmt/strconv afterwards
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: "use Error()",
						TextEdits: []analysis.TextEdit{{
							Pos:     node.Pos(),
							End:     call.Args[1].Pos(),
							NewText: []byte(""),
						},
							{
								Pos:     call.Args[1].End(),
								End:     node.End(),
								NewText: []byte(".Error()"),
							},
						},
					},
				},
			})
		} else if types.Identical(v, types.Typ[types.Bool]) && base == 1 {
			pass.Report(analysis.Diagnostic{
				Pos:     node.Pos(),
				End:     node.End(),
				Message: "fmt.Sprintf can be replaced with faster function strconv.FormatBool",
				// need ro run goimports to fix use of fmt/strconv afterwards
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: "use strconv.FormatBool",
						TextEdits: []analysis.TextEdit{{
							Pos:     node.Pos(),
							End:     call.Args[1].Pos(),
							NewText: []byte("strconv.FormatBool("),
						}},
					},
				},
			})
		} else if isarray && types.Identical(a.Elem(), types.Typ[types.Uint8]) && base == 16 {
			_, ok = call.Args[1].(*ast.Ident)
			if ok {
				pass.Report(analysis.Diagnostic{
					Pos:     node.Pos(),
					End:     node.End(),
					Message: "fmt.Sprintf can be replaced with faster function hex.EncodeToString",
					// need ro run goimports to fix use of fmt/encoding/hex afterwards
					SuggestedFixes: []analysis.SuggestedFix{
						{
							Message: "use hex.EncodeToString",
							TextEdits: []analysis.TextEdit{{
								Pos:     node.Pos(),
								End:     call.Args[1].Pos(),
								NewText: []byte("hex.EncodeToString("),
							},
								{
									Pos:     call.Args[1].End(),
									End:     call.Args[1].End(),
									NewText: []byte("[:]"),
								},
							},
						},
					},
				})
			}
		} else if isslice && types.Identical(s.Elem(), types.Typ[types.Uint8]) && base == 16 {
			pass.Report(analysis.Diagnostic{
				Pos:     node.Pos(),
				End:     node.End(),
				Message: "fmt.Sprintf can be replaced with faster function hex.EncodeToString",
				// need ro run goimports to fix use of fmt/strconv afterwards
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: "use hex.EncodeToString",
						TextEdits: []analysis.TextEdit{{
							Pos:     node.Pos(),
							End:     call.Args[1].Pos(),
							NewText: []byte("hex.EncodeToString("),
						}},
					},
				},
			})
		} else if types.Identical(v, types.Typ[types.Int]) && base == 10 {
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
		} else if types.Identical(v.Underlying(), types.Typ[types.Int]) && base == 10 {
		} else if types.Identical(v.Underlying(), types.Typ[types.Int8]) || types.Identical(v.Underlying(), types.Typ[types.Int16]) || types.Identical(v.Underlying(), types.Typ[types.Int32]) {
		} else if types.Identical(v.Underlying(), types.Typ[types.Uint8]) || types.Identical(v.Underlying(), types.Typ[types.Uint16]) || types.Identical(v.Underlying(), types.Typ[types.Uint32]) || types.Identical(v.Underlying(), types.Typ[types.Uint]){
		} else if arg0.Value == `"%v"` {
		} else if base == 0 {
		} else {
			pass.Reportf(node.Pos(), "Sprintf can be replaced with faster function from strconv %s/%s", arg0.Value, v.String())
		}
	})

	return nil, nil
}
