package service

import (
	domain "github.com/badcode256/example_go_hexagonal/internal/domain"
)

type UserService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) UserService {
	return UserService{
		userRepository: userRepository,
	}
}

func (s UserService) CreateUser(user domain.IUser) error {

	return s.userRepository.CreateUser(user)
}

func (s UserService) UpdateUser(user domain.UUser) error {

	return s.userRepository.UpdateUser(user)
}

func (s UserService) DeleteUser(id int64) error {

	return s.userRepository.DeleteUser(id)
}

func (s UserService) ListUsers() (*[]domain.Users, error) {

	return s.userRepository.ListUsers()
}
