package internal

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"
	"text/template"
)

type GenStruct struct {
	Name    string
	Fields  []GenParam
	Methods []GenMethod
}

type GenMethod struct {
	RecvName     string
	StructName   string
	Name         string
	Params       []GenParam
	Res          []GenParam
	ReturnFields []GenReturnField
}

type GenParam struct {
	Name string
	Type string
}

type GenReturnField struct {
	Recv string
	Name string
}

func createGenStructFromInterface(parsedInterface Interface) GenStruct {
	structName := "Stub" + parsedInterface.Name
	recv := "s"
	structFields := make([]GenParam, 0)
	methods := make([]GenMethod, 0)
	for _, parsedMethod := range parsedInterface.Methods {
		params := make([]GenParam, 0)
		for _, parsedParams := range parsedMethod.Params {
			paramName := parsedParams.Name
			if paramName == "" {
				paramName = "_"
			}
			params = append(params, GenParam{
				Name: paramName,
				Type: parsedParams.Type,
			})
		}
		res := make([]GenParam, 0)
		returnFields := make([]GenReturnField, 0)
		for i, parsedRes := range parsedMethod.Res {
			returnFieldName := fmt.Sprintf("%sRes%d", parsedMethod.Name, i)
			structFields = append(structFields, GenParam{
				Name: returnFieldName,
				Type: parsedRes.Type,
			})
			res = append(res, GenParam{
				Name: parsedRes.Name,
				Type: parsedRes.Type,
			})
			returnFields = append(returnFields, GenReturnField{
				Recv: recv,
				Name: returnFieldName,
			})
		}
		methods = append(methods, GenMethod{
			RecvName:     recv,
			StructName:   structName,
			Name:         parsedMethod.Name,
			Params:       params,
			Res:          res,
			ReturnFields: returnFields,
		})
	}
	return GenStruct{
		Name:    structName,
		Fields:  structFields,
		Methods: methods,
	}
}

func genStubSrcFromStruct(genStruct GenStruct) []byte {
	buf := new(bytes.Buffer)
	err := stubTemplate.Execute(buf, genStruct)
	if err != nil {
		panic(fmt.Errorf("error filling template: %w", err))
	}
	pretty, err := format.Source(buf.Bytes())
	if err != nil {
		panic(fmt.Errorf("error formatting stubs code: %w", err))
	}
	return pretty
}

func GenStubFromInterface(parsedInterface Interface) string {
	genStruct := createGenStructFromInterface(parsedInterface)
	stubSrc := genStubSrcFromStruct(genStruct)
	return strings.Trim(string(stubSrc), "\n")
}

const stubTemplateString = `
type {{.Name}} struct {
	{{range .Fields}} {{.Name}} {{.Type}}
	{{end}}

}
{{range .Methods}} 
func ({{.RecvName}} {{.StructName}}) {{.Name}} ({{range .Params}}{{.Name}} {{.Type}}, {{end}}) ({{range .Res}}{{.Name}} {{.Type}}, {{end}}){
	return {{range $i, $returnField := .ReturnFields}} {{if $i}},{{end}} {{.Recv}}.{{.Name}}{{end}}
}
{{end}}
`

var stubTemplate = template.Must(template.New("stub template").Parse(stubTemplateString))
