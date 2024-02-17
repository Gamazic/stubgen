package internal

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type AstModule struct {
	File           *ast.File
	Imports        []*ast.ImportSpec
	InterfaceTypes []*ast.TypeSpec
}

func GetAstModule(filename string, src []byte) (AstModule, error) {
	f, err := parser.ParseFile(token.NewFileSet(), filename, src, parser.SkipObjectResolution)
	if err != nil {
		return AstModule{}, fmt.Errorf("parsing source file as ast: %w", err)
	}
	foundInterfaces := make([]*ast.TypeSpec, 0)
	foundImports := make([]*ast.ImportSpec, 0)

	// depth-first search for interfaces and imports
	ast.Inspect(f, func(n ast.Node) bool {
		switch t := n.(type) {
		// find variable declaration
		case *ast.TypeSpec:
			itype, ok := t.Type.(*ast.InterfaceType)
			if ok && itype.Methods.List != nil {
				foundInterfaces = append(foundInterfaces, t)
			}
		case *ast.ImportSpec:
			foundImports = append(foundImports, t)
		}
		return true
	})
	return AstModule{
		File:           f,
		Imports:        foundImports,
		InterfaceTypes: foundInterfaces,
	}, nil
}
