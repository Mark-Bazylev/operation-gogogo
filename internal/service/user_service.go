package service

import (
	"operation-gogogo/internal/domain"
	"operation-gogogo/internal/repository"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (service *UserService) GetAllUsers() ([]domain.User, error) {
	return service.UserRepository.GetAllUsers()
}
