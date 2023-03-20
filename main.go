package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"log"
	"os"

	"golang.org/x/tools/go/ast/astutil"
)

func changeCall(call *ast.CallExpr) {
	name := call.Fun.(*ast.SelectorExpr)
	pname := name.X.(*ast.Ident)
	pname.Name = "strconv"
	name.Sel.Name = "Itoa"
	call.Args = call.Args[1:2]
	tocast := true
	switch a := call.Args[0].(type) {
	case *ast.CallExpr:
		if len(a.Args) == 1 {
			fn, ok := a.Fun.(*ast.Ident)
			if ok {
				switch fn.Name {
				case "int":
					tocast = false
				case "int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64":
					tocast = false
					fn.Name = "int"
				}
			}
		}
	}
	if tocast {
		call.Args[0] = &ast.CallExpr{Fun: &ast.Ident{Name: "int"}, Args: []ast.Expr{call.Args[0]}}
	}
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatalf("Expects a golang file")
	}
	path := flag.Args()[0]

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open golang file")
	}
	src, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read golang file")
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, src, parser.ParseComments)
	if err != nil {
		log.Fatalf("Failed to parse golang file")
	}
	changed := false
	ast.Inspect(f, func(n ast.Node) bool {
		switch call := n.(type) {
		case *ast.CallExpr:
			switch name := call.Fun.(type) {
			case *ast.SelectorExpr:
				if fmt.Sprintf("%s.%s", name.X, name.Sel) == "fmt.Sprintf" {
					if len(call.Args) == 2 {
						switch arg := call.Args[0].(type) {
						case *ast.BasicLit:
							if arg.Value == `"%d"` {
								changed = true
								changeCall(call)
							}
						}
					}
				}
			}
		}
		return true
	})
	if changed {
		if !astutil.UsesImport(f, "fmt") {
			astutil.DeleteImport(fset, f, "fmt")
		}
		astutil.AddImport(fset, f, "strconv")
		printer.Fprint(os.Stdout, fset, f)
	}
}
