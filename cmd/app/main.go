package main

import (
	"github.com/labstack/echo/v4" // Импортируем библиотеку Echo для работы с HTTP-запросами.
	"github.com/labstack/echo/v4/middleware"
	"log"
	"pet_project_1_etap/internal/database"    // Импортируем пакет для работы с базой данных.
	"pet_project_1_etap/internal/handlers"    // Импортируем пакет с обработчиками HTTP-запросов.
	"pet_project_1_etap/internal/taskService" // Импортируем пакет с сервисом задач.
	"pet_project_1_etap/internal/web/tasks"
)

func main() {
	// Инициализация базы данных
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	// Создаем репозиторий и сервис
	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	// Создаем базовый хендлер
	handler := handlers.NewHandler(service)

	// Инициализируем Echo
	e := echo.New()

	// Используем middleware для логирования и восстановления после паник
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Создаем строгий хендлер для интеграции с сгенерированным API
	strictHandler := tasks.NewStrictHandler(handler, nil) // Замените nil, если необходимо передать дополнительные middlewares

	// Регистрируем хендлеры с помощью сгенерированного RegisterHandlers
	tasks.RegisterHandlers(e, strictHandler)

	// Запускаем сервер
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
