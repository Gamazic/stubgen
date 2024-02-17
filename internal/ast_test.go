package internal

import (
	"testing"
)

func TestGetAstInterfaces(t *testing.T) {
	type args struct {
		filename string
		src      []byte
	}
	tests := []struct {
		name                  string
		args                  args
		wantFile              bool
		wantLenImports        int
		wantLenInterfaceTypes int
		wantErr               bool
	}{
		{
			name: "empty",
			args: args{
				filename: "",
				src:      []byte(""),
			},
			wantLenInterfaceTypes: 0,
			wantLenImports:        0,
			wantFile:              false,
			wantErr:               true,
		},
		{
			name: "package only",
			args: args{
				filename: "",
				src:      []byte("package main"),
			},
			wantLenInterfaceTypes: 0,
			wantLenImports:        0,
			wantFile:              true,
			wantErr:               false,
		},
		{
			name: "package, 1 import",
			args: args{
				filename: "",
				src: []byte(`
package name

import (
	e "internal"
)
`),
			},
			wantLenInterfaceTypes: 0,
			wantLenImports:        1,
			wantFile:              true,
			wantErr:               false,
		},
		{
			name: "1 interface, 1 import",
			args: args{
				filename: "",
				src: []byte(`
package testdata

import (
	"internal"
)

type MyInterface interface {
	MyFunc(arg int, arg2 bool) (ok bool, err error)
	Func2(int) error
}
`),
			},
			wantLenInterfaceTypes: 1,
			wantLenImports:        1,
			wantFile:              true,
			wantErr:               false,
		},
		{
			name: "1 interface, 2 import",
			args: args{
				filename: "",
				src: []byte(`
package testdata

import (
	"internal"
    e "path/to/external"
)

type MyInterface interface {
	MyFunc(arg int, arg2 e.Struct) (ok bool, err error)
	Func2(int) error
}
`),
			},
			wantLenInterfaceTypes: 1,
			wantLenImports:        2,
			wantFile:              true,
			wantErr:               false,
		},
		{
			name: "2 interfaces",
			args: args{
				filename: "",
				src: []byte(`
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
`),
			},
			wantLenInterfaceTypes: 2,
			wantLenImports:        0,
			wantFile:              true,
			wantErr:               false,
		},
		{
			name: "2 interfaces, 1 empty",
			args: args{
				filename: "",
				src: []byte(`
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
`),
			},
			wantLenInterfaceTypes: 2,
			wantLenImports:        0,
			wantFile:              true,
			wantErr:               false,
		},
		{
			name: "2 empty interfaces",
			args: args{
				filename: "",
				src: []byte(`
package testdata

type EmptyInterface1 interface{}
type EmptyInterface2 interface{}
`),
			},
			wantLenInterfaceTypes: 0,
			wantLenImports:        0,
			wantFile:              true,
			wantErr:               false,
		},

		{
			name: "2 empty interfaces with no package",
			args: args{
				filename: "",
				src: []byte(`
type EmptyInterface1 interface{}
type EmptyInterface2 interface{}
`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAstModule(tt.args.filename, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAstModule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got.File != nil) != tt.wantFile {
				t.Errorf("GetAstModule() .File = %v, wantFile %v", got.File, tt.wantErr)
				return
			}
			if len(got.Imports) != tt.wantLenImports {
				t.Errorf("GetAstModule() len(got.Imports) = %v, wantLenImports %v", got, tt.wantLenInterfaceTypes)
			}
			if len(got.InterfaceTypes) != tt.wantLenInterfaceTypes {
				t.Errorf("GetAstModule() len(got.InterfaceTypes) = %v, wantLenInterfaceTypes %v", got, tt.wantLenInterfaceTypes)
			}
		})
	}
}
