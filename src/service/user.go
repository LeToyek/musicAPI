package service

import "musicAPI/src/entities"

type UserService interface {
	AddUser(user entities.User)
	GetUser(email string, password string) (entities.User, error)
	DeleteUser(user entities.User) error
	EditUser(user entities.User, newUser entities.User) error
}

func (s *Service) AddUser(user entities.User) {
	s.Repository.AddUser(user)
}
func (s *Service) GetUser(email string, password string) (entities.User, error) {
	return s.Repository.GetUser(email, password)
}
func (s *Service) DeleteUser(user entities.User) error {
	return s.Repository.DeleteUser(user)
}
func (s *Service) EditUser(user entities.User, newUser entities.User) error {
	return s.Repository.EditUser(user, newUser)
}
