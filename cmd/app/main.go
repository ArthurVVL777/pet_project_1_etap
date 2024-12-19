package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"pet_project_1_etap/internal/database"
	"pet_project_1_etap/internal/handlers"
	"pet_project_1_etap/internal/taskService"
	"pet_project_1_etap/internal/web/tasks"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Добавляем обработку ошибок при миграции
	if err := database.DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрация обработчиков с учетом строгой сигнатуры
	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	// Запуск сервера на порту 8080
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
