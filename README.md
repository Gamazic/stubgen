# stubgen

A Golang library for generating interface stubs.

## Installation

```shell
go install github.com/Gamazic/stubgen@latest
```

## Example

Have go module with the following content: 

`testfile.go`
```go
package testdata

type MyInterface interface {
	Func(int, bool) error
}

```

To generate stub for the file use the command with arg `--inp-file`:

```shell
stubgen --inp-file testdata/testfile.go
```

Result:

```go
package testdata

type StubMyInterface struct {
	FuncRes0 error
}

func (s StubMyInterface) Func(_ int, _ bool) error {
	return s.FuncRes0
}

```

Alternatively, you can provide a source code from stdin:

```shell
cat testdata/testfile.go | stubgen
```

To save data in a file use the `--out-file` arg:

```shell
stubgen --inp-file testfile.go --out-file testfile_stub.go
```
