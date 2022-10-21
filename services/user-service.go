package services

import (
	"errors"

	"github.com/titusdishon/go-docker-mysql/entity"
	"github.com/titusdishon/go-docker-mysql/repositories"
)

var (
	repo repositories.UserRepository
)

type UserService interface {
	Validate(user *entity.User) error
	Save(user *entity.User) (*entity.User, error)
	Update(user *entity.User, id int64) (*entity.User, error)
	FindAll() ([]entity.User, error)
	FindById(id int64) (*entity.User, error)
	Delete(id int64) (int64, error)
}
type service struct{}

func NewUserService(repository repositories.UserRepository) UserService {
	repo = repository
	return &service{}
}

func (*service) Validate(user *entity.User) error {
	if user == nil {
		err := errors.New("user object is empty")
		return err
	}
	if user.Name == "" {
		err := errors.New("name field is empty")
		return err
	}
	if user.Email == "" {
		err := errors.New("email field is empty")
		return err
	}
	if user.Summary == "" {
		err := errors.New("summary field is empty")
		return err
	}
	return nil
}
func (*service) Save(user *entity.User) (*entity.User, error) {
	return repo.Save(user)
}
func (*service) Update(user *entity.User, id int64) (*entity.User, error) {
	return repo.Update(user, id)
}
func (*service) FindAll() ([]entity.User, error) {
	return repo.FindAll()
}
func (*service) FindById(id int64) (*entity.User, error) {
	return repo.FindById(id)
}
func (*service) Delete(id int64) (int64, error) {
	return repo.Delete(id)
}
