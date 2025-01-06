package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
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
func (u *UserHandler) GetUsers(ctx echo.Context) error {
	user, err := u.Service.GetAllUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, user)
}

// PostUsers обрабатывает запрос на создание нового пользователя.
func (u *UserHandler) PostUsers(ctx echo.Context) error {
	var user userService.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	createdUser, err := u.Service.CreateUser(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, createdUser)
}

// PatchUserByID обрабатывает запрос на обновление существующего пользователя.
func (u *UserHandler) PatchUserByID(ctx echo.Context, id uint) error {
	var user userService.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	updatedUser, err := u.Service.UpdateUserByID(id, user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, updatedUser)
}

// DeleteUserByID обрабатывает запрос на удаление пользователя по ID.
func (u *UserHandler) DeleteUserByID(ctx echo.Context, id uint) error {
	if err := u.Service.DeleteUserByID(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}
