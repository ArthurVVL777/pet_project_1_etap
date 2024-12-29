package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"pet_project_1_etap/internal/database"
	"pet_project_1_etap/internal/handlers"
	"pet_project_1_etap/internal/taskService"
	"pet_project_1_etap/internal/userService"
	"pet_project_1_etap/internal/web/tasks"
	"pet_project_1_etap/internal/web/users"
)

func main() {
	database.InitDB()
	database.Migrate()

	// Обновление названий репозиториев и сервисов задач
	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksSvc := taskService.NewService(tasksRepo)
	tasksHandler := handlers.NewHandler(tasksSvc)

	// Добавление репозитория и сервиса пользователей
	userRepo := userService.NewUserRepository(database.DB)
	userSvc := userService.NewService(userRepo)
	userHandler := handlers.NewUserHandler(userSvc)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрация обработчиков задач и пользователей
	tasks.RegisterHandlers(e, tasksHandler)
	users.RegisterHandlers(e, userHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
