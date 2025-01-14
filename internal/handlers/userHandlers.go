package handlers

import (
	"fmt"
	"log"
	"net/http"
	"pet_project_1_etap/internal/userService"
	"pet_project_1_etap/internal/web/users"

	"github.com/labstack/echo/v4"
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
	var request users.PostUsersRequestObject
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// Проверка на nil для полей Email и Password
	if request.Body == nil || request.Body.Email == nil || request.Body.Password == nil {
		return ctx.JSON(http.StatusBadRequest, "Email and password must not be empty")
	}

	user := userService.User{
		Email:    *request.Body.Email,
		Password: *request.Body.Password,
	}

	createdUser, err := u.Service.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return ctx.JSON(http.StatusInternalServerError, "Error creating user")
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

	existingUser, err := u.Service.GetUserByID(id)
	if err != nil {
		log.Printf("Error fetching user with ID %d: %v", id, err)
		return ctx.JSON(http.StatusNotFound, fmt.Sprintf("user not found: %v", err))
	}

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
		log.Printf("Error deleting user with ID %d: %v", id, err)
		return ctx.JSON(http.StatusInternalServerError, "Error deleting user")
	}

	return ctx.NoContent(http.StatusNoContent)
}
