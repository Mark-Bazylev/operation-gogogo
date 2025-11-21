package repository

import "operation-gogogo/internal/domain"

type UserRepository interface {
	GetAllUsers() ([]domain.User, error)
}
