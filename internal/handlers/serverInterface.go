package handlers

import (
	"context"
	"pet_project_1_etap/internal/userService"
	"pet_project_1_etap/internal/web/tasks"
)

// ServerInterface определяет методы для работы с задачами и пользователями.
type ServerInterface interface {
	// Методы для работы с задачами
	GetTasks(ctx context.Context, req tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error)
	PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error)
	PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error)
	DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error)

	// Методы для работы с пользователями
	GetUsers(ctx context.Context) ([]userService.User, error)
	PostUser(ctx context.Context, user userService.User) (userService.User, error)
	PatchUserByID(ctx context.Context, id uint, user userService.User) (userService.User, error)
	DeleteUserByID(ctx context.Context, id uint) error
}
