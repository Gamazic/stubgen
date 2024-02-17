package internal

import (
	"fmt"
	"go/ast"
	"strings"
)

func ParsePackageName(f *ast.File) string {
	return f.Name.Name
}

// ParseImports returns {'packAlias': 'packAlias "path/to/pack"', 'packName': '"path/to/packName"'}
func ParseImports(astImports []*ast.ImportSpec) map[string]string {
	imports := make(map[string]string)
	for _, imp := range astImports {
		if imp.Name != nil { // if there is an alias
			imports[imp.Name.Name] = fmt.Sprintf("%s %s", imp.Name.Name, imp.Path.Value)
		} else {
			path := strings.Split(strings.Trim(imp.Path.Value, "\""), "/")
			packName := path[len(path)-1]
			imports[packName] = imp.Path.Value
		}
	}
	return imports
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
	Type Type
}

type Type struct {
	Package string
	Name    string
}

func (t Type) String() string {
	if t.Package == "" {
		return t.Name
	}
	return fmt.Sprintf("%s.%s", t.Package, t.Name)
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
		// string -> {Name: "", Type: "string"}
		// arg string -> {Name: "arg", Type: "string"}
		// arg1, arg2 string -> {Name: "arg1, arg2", Type: "string"}
		paramSameTypeNames := make([]string, 0)
		for _, astName := range astParam.Names {
			paramSameTypeNames = append(paramSameTypeNames, astName.Name)
		}
		paramName := strings.Join(paramSameTypeNames, ", ")
		params = append(params, Param{
			Name: paramName,
			Type: fullType(astParam),
		})
	}
	return params
}

func fullType(astParam *ast.Field) Type {
	switch astParam.Type.(type) {
	case *ast.Ident:
		return Type{
			Package: "",
			Name:    astParam.Type.(*ast.Ident).Name,
		}
	case *ast.SelectorExpr:
		typeExrp := astParam.Type.(*ast.SelectorExpr)
		packName := typeExrp.X.(*ast.Ident)
		return Type{
			Package: packName.Name,
			Name:    typeExrp.Sel.Name,
		}
	default:
		panic("not expected param ast type")
	}
}
