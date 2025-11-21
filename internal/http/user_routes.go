package http

import (
	"net/http"

	"operation-gogogo/internal/repository"
	"operation-gogogo/internal/service"
)

func registerUserRoutes(mux *http.ServeMux) {
	userRepository := &repository.MockUserRepository{}

	userService := &service.UserService{
		UserRepository: userRepository,
	}

	userHandler := &UserHandler{
		UserService: userService,
	}

	mux.HandleFunc("/v1/users", userHandler.GetAllUsers)
}
