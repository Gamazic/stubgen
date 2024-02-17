package example

type UseCase struct {
	Repo UserRepo
}

func (u UseCase) CreateUser(name string) error {
	userId, err := u.Repo.NextId()
	if err != nil {
		return err
	}
	err = u.Repo.Store(userId, name)
	if err != nil {
		return err
	}
	return nil
}

type UserRepo interface {
	NextId() (int, error)
	Store(id int, name string) error
}
