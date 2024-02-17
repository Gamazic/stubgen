package testdata

import "github.com/Gamazic/stubgen/testdata/internal"

type Interface7 interface {
	Method7() error
}

type Interface8 interface {
	M1(int, bool) error
	M2(arg1 int, arg2 bool) (ok bool, err error)
	M3(arg1, arg2 string) internal.Error
	Empty()
}
