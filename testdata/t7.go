package testdata

import (
	"errors"
	"github.com/Gamazic/stubgen/testdata/internal"
	"io"
)

type MyInterface interface {
	AnotherFunc(int, internal.MyStruct) error
}

type User struct {
	m MyInterface
	r io.Reader
}

func (User) Foo(a int) error {
	return errors.New("")
}
