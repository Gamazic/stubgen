package internal

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func GetAstInterfaces(filename, src string) ([]*ast.TypeSpec, error) {
	f, err := parser.ParseFile(token.NewFileSet(), filename, src, parser.SkipObjectResolution)
	if err != nil {
		return nil, fmt.Errorf("parsing source file as ast: %w", err)
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
	return foundInterfaces, nil
}
