package repository

import (
	"github.com/go-gorp/gorp"
	"github.com/kiki-ki/go-test-example/internal/app/model"
)

type executer = gorp.SqlExecutor

func NewUserRepository() UserRepository {
	return &userRepository{}
}

type UserRepository interface {
	All(e executer) ([]model.User, error)
	Find(uId int, e executer) (model.User, error)
	Update(u *model.User, e executer) error
	Create(u *model.User, e executer) error
	Delete(uId int, e executer) error
}

type userRepository struct{}

func (r *userRepository) All(e executer) ([]model.User, error) {
	var users []model.User
	_, err := e.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Find(uId int, e executer) (model.User, error) {
	var u model.User
	err := e.SelectOne(&u, "SELECT * FROM users WHERE id = ?", uId)
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (r *userRepository) Update(u *model.User, e executer) error {
	_, err := e.Update(u)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Create(u *model.User, e executer) error {
	err := e.Insert(u)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(uId int, e executer) error {
	_, err := e.Exec("delete from users where id=?", uId)
	if err != nil {
		return err
	}
	return nil
}
