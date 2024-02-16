package internal

import (
	"fmt"
	"go/ast"
)

func ParsePackage(f *ast.File) string {
	return f.Name.Name
}

type Interface struct {
	Name    string
	Methods []Method
}

type Method struct {
	Name   string
	Params []Param
	Res    []Param
}

type Param struct {
	Name string
	Type string
}

func ParseInterface(iType *ast.TypeSpec) Interface {
	astInterface := iType.Type.(*ast.InterfaceType)
	methods := make([]Method, astInterface.Methods.NumFields())
	for i, astMethod := range astInterface.Methods.List {
		astFunc := astMethod.Type.(*ast.FuncType)
		params := parseParams(astFunc.Params)
		res := parseParams(astFunc.Results)
		methods[i] = Method{
			Name:   astMethod.Names[0].String(),
			Params: params,
			Res:    res,
		}
	}
	return Interface{
		Name:    iType.Name.String(),
		Methods: methods,
	}
}

func parseParams(fields *ast.FieldList) []Param {
	if fields == nil {
		return nil
	}
	params := make([]Param, 0)
	for _, astParam := range fields.List {
		paramName := ""
		if astParam.Names != nil {
			paramName = astParam.Names[0].String()
		}
		params = append(params, Param{
			Name: paramName,
			Type: fullType(astParam),
		})
	}
	return params
}

func fullType(astParam *ast.Field) string {
	switch astParam.Type.(type) {
	case *ast.Ident:
		return astParam.Type.(*ast.Ident).Name
	case *ast.SelectorExpr:
		typeExrp := astParam.Type.(*ast.SelectorExpr)
		packName := typeExrp.X.(*ast.Ident)
		return fmt.Sprintf("%s.%s", packName.Name, typeExrp.Sel.Name)
	default:
		panic("not expected param ast type")
	}
}
