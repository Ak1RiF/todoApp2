package handlers

import (
	"github.com/todoApp/internal/repository"
	"github.com/todoApp/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler() *UserHandler {
	repo := repository.NewRepository()
	deps := service.Deps{Repositories: repo}
	return &UserHandler{
		userService: *service.NewUserService(deps.Repositories.Users),
	}
}

// handlers
