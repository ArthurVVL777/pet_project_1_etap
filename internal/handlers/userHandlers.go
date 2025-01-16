package handlers

import (
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
	var request users.User
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// Проверяем обязательные поля
	if request.Email == nil || *request.Email == "" || request.Password == nil || *request.Password == "" {
		return ctx.JSON(http.StatusBadRequest, "Email and Password must not be empty")
	}

	// Создаем нового пользователя
	createdUser, err := u.Service.CreateUser(userService.User{
		Email:    *request.Email,
		Password: *request.Password,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error creating user")
	}

	// Формируем ответ без ненужных полей
	response := users.User{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
		//Tasks:    createdUser.Tasks, // Возвращаем поле tasks без изменений
	}
	return ctx.JSON(http.StatusCreated, response)
}

func (u *UserHandler) PatchUsersId(ctx echo.Context, id uint) error {
	var request userService.PatchUserRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid request body")
	}

	existingUser, err := u.Service.GetUserByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "User not found")
	}

	// Обновляем только те поля, которые были указаны в запросе
	if request.Email != nil {
		existingUser.Email = *request.Email
	}
	if request.Password != nil {
		existingUser.Password = *request.Password
	}

	updatedUser, err := u.Service.UpdateUserByID(id, existingUser)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error updating user")
	}

	// Удаляем ненужные поля из ответа
	response := users.User{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
		//Tasks:    updatedUser.Tasks, // Поле tasks возвращаем, как есть
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
