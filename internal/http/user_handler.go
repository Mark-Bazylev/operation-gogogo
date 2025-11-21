package http

import (
	"encoding/json"
	"net/http"
	"operation-gogogo/internal/service"
)

type UserHandler struct {
	UserService *service.UserService
}

// GetAllUsers godoc
// @Summary      Get all users
// @Tags         users
// @Produce      json
// @Success      200  {array}  domain.User
// @Router       /v1/users [get]
func (handler *UserHandler) GetAllUsers(responseWriter http.ResponseWriter, request *http.Request) {
	users, serviceError := handler.UserService.GetAllUsers()
	if serviceError != nil {
		http.Error(responseWriter, "failed to fetch users", http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	encodeError := json.NewEncoder(responseWriter).Encode(users)
	if encodeError != nil {
		http.Error(responseWriter, "failed to encode users", http.StatusInternalServerError)
		return
	}
}
