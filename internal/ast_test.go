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
		name                  string
		args                  args
		wantLenInterfaceTypes int
		wantFile              bool
		wantErr               bool
	}{
		{
			name: "empty",
			args: args{
				filename: "",
				src:      "",
			},
			wantLenInterfaceTypes: 0,
			wantFile:              false,
			wantErr:               true,
		},
		{
			name: "package only",
			args: args{
				filename: "",
				src:      "package main",
			},
			wantLenInterfaceTypes: 0,
			wantFile:              true,
			wantErr:               false,
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
			wantLenInterfaceTypes: 1,
			wantFile:              true,
			wantErr:               false,
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
			wantLenInterfaceTypes: 2,
			wantFile:              true,
			wantErr:               false,
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
			wantLenInterfaceTypes: 2,
			wantFile:              true,
			wantErr:               false,
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
			wantLenInterfaceTypes: 0,
			wantFile:              true,
			wantErr:               false,
		},
		{
			name: "2 empty interfaces with no package",
			args: args{
				filename: "",
				src: `
type EmptyInterface1 interface{}
type EmptyInterface2 interface{}
`,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAstInfo(tt.args.filename, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAstInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got.File != nil) != tt.wantFile {
				t.Errorf("GetAstInfo() .File = %v, wantFile %v", got.File, tt.wantErr)
				return
			}
			if len(got.InterfaceTypes) != tt.wantLenInterfaceTypes {
				t.Errorf("GetAstInfo() len(got) = %v, wantLenInterfaceTypes %v", got, tt.wantLenInterfaceTypes)
			}
		})
	}
}
