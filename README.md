# stubgen

A Golang library for generating interface stubs.

## Installation

```shell
go install github.com/Gamazic/stubgen@latest
```

## Example

For example, we have go module with the following content: 

`testfile.go`
```go
package testdata

type MyInterface interface {
	AnotherFunc(int, bool) error
}

```

To generate stub for the file use command with arg `--inp-file`:

```shell
stubgen --inp-file testdata/testfile.go
```

Result:

```go
package testdata

type StubMyInterface struct {
	AnotherFuncRes0 error
}

func (s StubMyInterface) AnotherFunc(_ int, _ bool) error {
	return s.AnotherFuncRes0
}

```

Alternatively, you can provide source code from stdin:

```shell
cat testdata/testfile.go | stubgen
```

To save data in a file use the `--out-file` arg:

```shell
stubgen --inp-file testdata/testfile.go --out-file interface.go
```
