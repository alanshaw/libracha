package main

import (
	"fmt"
	"go/ast"
	"os"

	"golang.org/x/tools/go/packages"
)

const marker = "// +libracha:map-encoders"

func main() {
	patterns := os.Args[1:] // e.g. ["./..."]
	if len(patterns) == 0 {
		patterns = []string{"."}
	}

	cfg := &packages.Config{Mode: packages.NeedSyntax | packages.NeedTypes}
	pkgs, _ := packages.Load(cfg, patterns...)

	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			for _, decl := range file.Decls {
				genDecl, ok := decl.(*ast.GenDecl)
				if !ok {
					continue
				}
				if genDecl.Doc == nil {
					continue
				}
				for _, comment := range genDecl.Doc.List {
					if comment.Text == marker {
						for _, spec := range genDecl.Specs {
							typeSpec, ok := spec.(*ast.TypeSpec)
							if !ok {
								continue
							}
							structType, ok := typeSpec.Type.(*ast.StructType)
							if !ok {
								continue
							}
							fmt.Println("Found struct:", typeSpec.Name.Name)
							for _, field := range structType.Fields.List {
								for _, name := range field.Names {
									fmt.Println("  Field:", name.Name)
								}
							}
						}
					}
				}
			}
		}
	}
}
