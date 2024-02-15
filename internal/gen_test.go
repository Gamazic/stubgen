package internal

import "testing"

func TestGenStubFromInterface(t *testing.T) {
	type args struct {
		parsedInterface Interface
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				parsedInterface: Interface{
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
			want: stubExample1,
		},
		{
			name: "",
			args: args{
				parsedInterface: Interface{
					Name: "MyInterface",
					Methods: []Method{
						{
							Name: "AnotherFunc",
							Params: []Param{
								{
									Name: "",
									Type: "int",
								},
								{
									Name: "",
									Type: "bool",
								},
							},
							Res: []Param{
								{
									Name: "",
									Type: "error",
								},
							},
						},
					},
				},
			},
			want: stubExample2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenStubFromInterface(tt.args.parsedInterface); got != tt.want {
				t.Errorf("GenStubFromInterface() = \n%v\n, want \n%v", got, tt.want)
			}
		})
	}
}

const (
	stubExample1 = "" +
		`type StubMyInterface struct {
	MyFuncRes0 bool
	MyFuncRes1 external.MyStruct2
	MyFuncRes2 error
}

func (s StubMyInterface) MyFunc(arg int, arg2 internal.MyStruct1) (ok bool, s external.MyStruct2, err error) {
	return s.MyFuncRes0, s.MyFuncRes1, s.MyFuncRes2
}`
	stubExample2 = "" +
		`type StubMyInterface struct {
	AnotherFuncRes0 error
}

func (s StubMyInterface) AnotherFunc(_ int, _ bool) error {
	return s.AnotherFuncRes0
}`
)
