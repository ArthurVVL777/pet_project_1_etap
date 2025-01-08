package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"pet_project_1_etap/internal/userService"
	"pet_project_1_etap/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

// NewUserHandler создает новый экземпляр UserHandler с заданным сервисом.
func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) GetUsers(ctx echo.Context) error {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return ctx.JSON(http.StatusInternalServerError, "Error fetching users")
	}

	response := make([]users.User, 0, len(allUsers))
	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	return ctx.JSON(http.StatusOK, response)
}

func (u *UserHandler) PostUsers(ctx echo.Context) error {
	var user userService.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	createdUser, err := u.Service.CreateUser(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	response := users.User{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return ctx.JSON(http.StatusCreated, response)
}

func (u *UserHandler) PatchUsersId(ctx echo.Context, id uint) error {
	var request users.PatchUserIdRequestObject
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}

	existingUser, err := u.Service.GetUserByID(id) // Используем id напрямую
	if err != nil {
		log.Printf("Error fetching user with ID %d: %v", id, err)
		return ctx.JSON(http.StatusNotFound, fmt.Sprintf("user not found: %v", err))
	}

	// Обновляем только те поля, которые были переданы
	if request.Body.Email != nil {
		existingUser.Email = *request.Body.Email
	}

	if request.Body.Password != nil {
		existingUser.Password = *request.Body.Password
	}

	updatedUser, err := u.Service.UpdateUserByID(id, existingUser)
	if err != nil {
		log.Printf("Error updating user with ID %d: %v", id, err)
		return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to update user: %v", err))
	}

	response := users.User{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (u *UserHandler) DeleteUsersId(ctx echo.Context, id uint) error {
	if err := u.Service.DeleteUserByID(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}
