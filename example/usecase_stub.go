// Code generated by stubgen; DO NOT EDIT.

package example

type StubUserRepo struct {
	NextIdRes0 int
	NextIdRes1 error
	StoreRes0  error
}

func (s StubUserRepo) NextId() (int, error) {
	return s.NextIdRes0, s.NextIdRes1
}

func (s StubUserRepo) Store(id int, name string) error {
	return s.StoreRes0
}
