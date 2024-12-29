package handlers

import (
	"context"
	"pet_project_1_etap/internal/userService"
)

type UserHandler struct {
	Service *userService.UserService
}

// NewUserHandler создает новый экземпляр UserHandler с заданным сервисом.
func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// GetUsers обрабатывает запрос на получение всех пользователей.
func (h *UserHandler) GetUsers(ctx context.Context) ([]userService.User, error) {
	return h.Service.GetAllUsers()
}

// PostUser обрабатывает запрос на создание нового пользователя.
func (h *UserHandler) PostUser(ctx context.Context, user userService.User) (userService.User, error) {
	return h.Service.CreateUser(user)
}

// PatchUserByID обрабатывает запрос на обновление существующего пользователя.
func (h *UserHandler) PatchUserByID(ctx context.Context, id uint, user userService.User) (userService.User, error) {
	return h.Service.UpdateUserByID(id, user)
}

// DeleteUserByID обрабатывает запрос на удаление пользователя по ID.
func (h *UserHandler) DeleteUserByID(ctx context.Context, id uint) error {
	return h.Service.DeleteUserByID(id)
}
