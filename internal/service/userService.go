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
