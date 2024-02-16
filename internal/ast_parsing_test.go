package internal

import (
	"go/ast"
	"reflect"
	"testing"
)

func TestParseInterface(t *testing.T) {
	type args struct {
		iType *ast.TypeSpec
	}
	tests := []struct {
		name string
		args args
		want Interface
	}{
		{
			name: "",
			args: args{
				iType: &ast.TypeSpec{
					Name: &ast.Ident{
						Name: "MyInterface",
					},
					Type: &ast.InterfaceType{
						Methods: &ast.FieldList{
							List: []*ast.Field{
								{
									Names: []*ast.Ident{
										{Name: "MyFunc"},
									},
									Type: &ast.FuncType{
										Params: &ast.FieldList{
											List: []*ast.Field{
												{
													Names: []*ast.Ident{
														{
															Name: "arg",
														},
													},
													Type: &ast.Ident{
														Name: "int",
													},
												},
												{
													Names: []*ast.Ident{
														{
															Name: "arg2",
														},
													},
													Type: &ast.SelectorExpr{
														X: &ast.Ident{
															Name: "internal",
														},
														Sel: &ast.Ident{
															Name: "MyStruct1",
														},
													},
												},
											},
										},
										Results: &ast.FieldList{
											List: []*ast.Field{
												{
													Names: []*ast.Ident{
														{
															Name: "ok",
														},
													},
													Type: &ast.Ident{
														Name: "bool",
													},
												},
												{
													Names: []*ast.Ident{
														{
															Name: "s",
														},
													},
													Type: &ast.SelectorExpr{
														X: &ast.Ident{
															Name: "external",
														},
														Sel: &ast.Ident{
															Name: "MyStruct2",
														},
													},
												},
												{
													Names: []*ast.Ident{
														{
															Name: "err",
														},
													},
													Type: &ast.Ident{
														Name: "error",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: Interface{
				Name: "MyInterface",
				Methods: []Method{
					{
						Name: "MyFunc",
						Params: []Param{
							{
								Name: "arg",
								Type: "int",
							},
							{
								Name: "arg2",
								Type: "internal.MyStruct1",
							},
						},
						Res: []Param{
							{
								Name: "ok",
								Type: "bool",
							},
							{
								Name: "s",
								Type: "external.MyStruct2",
							},
							{
								Name: "err",
								Type: "error",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseInterface(tt.args.iType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}
