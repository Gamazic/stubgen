package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_genStubsFromSrc(t *testing.T) {
	mustReadFile := func(fname string) []byte {
		src, _ := os.ReadFile(fname)
		return src
	}

	type args struct {
		src []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "1 interface",
			args: args{
				src: mustReadFile("testdata/t1.go"),
			},
			want:    mustReadFile("testdata/t1_stub.go"),
			wantErr: false,
		},
		{
			name: "1 interface no args",
			args: args{
				src: mustReadFile("testdata/t2.go"),
			},
			want:    mustReadFile("testdata/t2_stub.go"),
			wantErr: false,
		},
		{
			name: "2 interfaces",
			args: args{
				src: mustReadFile("testdata/t3.go"),
			},
			want:    mustReadFile("testdata/t3_stub.go"),
			wantErr: false,
		},
		{
			name: "2 interfaces, one empty",
			args: args{
				src: mustReadFile("testdata/t4.go"),
			},
			want:    mustReadFile("testdata/t4_stub.go"),
			wantErr: false,
		},
		{
			name: "2 empty interface",
			args: args{
				src: mustReadFile("testdata/t5.go"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "big interfaces with import",
			args: args{
				src: mustReadFile("testdata/t6.go"),
			},
			want:    mustReadFile("testdata/t6_stub.go"),
			wantErr: false,
		},
		{
			name: "interface with other code",
			args: args{
				src: mustReadFile("testdata/t7.go"),
			},
			want:    mustReadFile("testdata/t7_stub.go"),
			wantErr: false,
		},
		{
			name: "short import",
			args: args{
				src: mustReadFile("testdata/t8.go"),
			},
			want:    mustReadFile("testdata/t8_stub.go"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := genStubsFromSrc(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("genStubsFromSrc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genStubsFromSrc() got = \n%v\n, want ----------"+
					" \n%v", string(got), string(tt.want))
			}
		})
	}
}
