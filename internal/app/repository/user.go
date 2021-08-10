package repository

import (
	"github.com/kiki-ki/go-test-example/internal/app/model"
	"github.com/kiki-ki/go-test-example/internal/interface/database"
)

func NewUserRepository(db database.DB) UserRepository {
	return &userRepository{db}
}

type UserRepository interface {
	All() ([]model.User, error)
	Find(uId int) (model.User, error)
	Update(u *model.User) error
	Create(u *model.User) error
	Delete(uId int) error
}

type userRepository struct {
	db database.DB
}

func (r *userRepository) All() ([]model.User, error) {
	var users []model.User
	_, err := r.db.Conn().Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Find(uId int) (model.User, error) {
	var u model.User
	err := r.db.Conn().SelectOne(&u, "SELECT * FROM users WHERE id = ?", uId)
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (r *userRepository) Update(u *model.User) error {
	_, err := r.db.Conn().Update(u)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Create(u *model.User) error {
	err := r.db.Conn().Insert(u)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(uId int) error {
	_, err := r.db.Conn().Exec("delete from users where id=?", uId)
	if err != nil {
		return err
	}
	return nil
}
