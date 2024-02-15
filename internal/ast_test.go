package internal

import (
	"testing"
)

func TestGetAstInterfaces(t *testing.T) {
	type args struct {
		filename string
		src      string
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
		wantErr bool
	}{
		{
			name: "empty",
			args: args{
				filename: "",
				src:      "",
			},
			wantLen: 0,
			wantErr: true,
		},
		{
			name: "package only",
			args: args{
				filename: "",
				src:      "package main",
			},
			wantLen: 0,
			wantErr: false,
		},
		{
			name: "one interface",
			args: args{
				filename: "",
				src: `
package testdata

type MyInterface interface {
	MyFunc(arg int, arg2 bool) (ok bool, err error)
	Func2(int) error
}
`,
			},
			wantLen: 1,
			wantErr: false,
		},
		{
			name: "2 interfaces",
			args: args{
				filename: "",
				src: `
package testdata

type MyInterface interface {
	MyFunc(arg int, arg2 bool) (ok bool, err error)
	Func2(int) error
}

// Interface2 COmment
type Interface2 interface {
	// Foo comment
	Foo() bool
}
`,
			},
			wantLen: 2,
			wantErr: false,
		},
		{
			name: "2 interfaces, 1 empty",
			args: args{
				filename: "",
				src: `
package testdata

type MyInterface interface {
	MyFunc(arg int, arg2 bool) (ok bool, err error)
	Func2(int) error
}

// Interface2 COmment
type Interface2 interface {
	// Foo comment
	Foo() bool
}

type EmptyInterface interface{}
`,
			},
			wantLen: 2,
			wantErr: false,
		},
		{
			name: "2 empty interfaces",
			args: args{
				filename: "",
				src: `
package testdata

type EmptyInterface1 interface{}
type EmptyInterface2 interface{}
`,
			},
			wantLen: 0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAstInterfaces(tt.args.filename, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAstInterfaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("GetAstInterfaces() len(got) = %v, wantLen %v", got, tt.wantLen)
			}
		})
	}
}
