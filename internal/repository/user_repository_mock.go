package repository

import "operation-gogogo/internal/domain"

type MockUserRepository struct{}

func (repository *MockUserRepository) GetAllUsers() ([]domain.User, error) {
	mockUsers := []domain.User{
		{
			ID:        1,
			FirstName: "Mark",
			LastName:  "Bazylev",
			Email:     "mark@example.com",
		},
		{
			ID:        2,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
		},
	}

	return mockUsers, nil
}
