package internal

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type AstInfo struct {
	File           *ast.File
	InterfaceTypes []*ast.TypeSpec
}

func GetAstInfo(filename, src string) (AstInfo, error) {
	f, err := parser.ParseFile(token.NewFileSet(), filename, src, parser.SkipObjectResolution)
	if err != nil {
		return AstInfo{}, fmt.Errorf("parsing source file as ast: %w", err)
	}
	foundInterfaces := make([]*ast.TypeSpec, 0)

	// depth-first search for interfaces
	ast.Inspect(f, func(n ast.Node) bool {
		switch t := n.(type) {
		// find variable declaration
		case *ast.TypeSpec:
			if t.Name.IsExported() {
				itype, ok := t.Type.(*ast.InterfaceType)
				if ok && itype.Methods.List != nil {
					foundInterfaces = append(foundInterfaces, t)
				}
			}
		}
		return true
	})
	return AstInfo{
		File:           f,
		InterfaceTypes: foundInterfaces,
	}, nil
}
