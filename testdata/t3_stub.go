// Code generated by stubgen; DO NOT EDIT.

package testdata

type StubInterface3 struct {
	Method3Res0 error
}

func (s StubInterface3) Method3() error {
	return s.Method3Res0
}

type StubInterface4 struct {
	Method4Res0 error
}

func (s StubInterface4) Method4(_ error) error {
	return s.Method4Res0
}
