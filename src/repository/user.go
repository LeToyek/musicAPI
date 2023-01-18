package repository

import (
	"errors"
	"musicAPI/src/entities"
	"musicAPI/src/repository/query"
)

type UserRepository interface {
	AddUser(user entities.User)
	GetUser(email string, password string) entities.User
	DeleteUser(user entities.User) error
	EditUser(user entities.User) error
}

func (r *Repository) AddUser(user entities.User) {
	r.DB.MustExec(query.QAddUser, user.ID, user.Username, user.Email, user.Password, user.Refresh_Token, user.Created_at, user.Updated_at)
}
func (r *Repository) GetUser(email string, password string) (entities.User, error) {
	result := entities.User{}
	err := r.DB.Get(&result, query.QGetUser, email, password)
	return result, err
}

func (r *Repository) DeleteUser(user entities.User) error {
	_, err := r.DB.NamedExec(query.QDeleteUser, user)
	return err
}
func (r *Repository) EditUser(user entities.User, newUser entities.User) error {
	return errors.New("unsupported feature, not finished yet")
}
