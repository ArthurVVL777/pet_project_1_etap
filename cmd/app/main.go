package main

import (
	"github.com/labstack/echo/v4"             // Импортируем библиотеку Echo для работы с HTTP-запросами.
	"pet_project_1_etap/internal/database"    // Импортируем пакет для работы с базой данных.
	"pet_project_1_etap/internal/handlers"    // Импортируем пакет с обработчиками HTTP-запросов.
	"pet_project_1_etap/internal/taskService" // Импортируем пакет с сервисом задач.
)

func main() {
	// Инициализация базы данных
	database.InitDB()                            // Вызываем функцию InitDB для установки соединения с базой данных.
	database.DB.AutoMigrate(&taskService.Task{}) // Автоматически применяем миграции для структуры Task, создавая необходимые таблицы в БД.

	// Создание репозитория и сервиса
	repo := taskService.NewTaskRepository(database.DB) // Создаем новый экземпляр репозитория задач, передавая подключение к базе данных.
	service := taskService.NewService(repo)            // Создаем новый экземпляр сервиса задач, передавая созданный репозиторий.

	// Создание обработчиков
	handler := handlers.NewHandler(service) // Создаем новый экземпляр обработчиков, передавая созданный сервис задач.

	// Создание нового экземпляра Echo
	e := echo.New() // Инициализируем новый экземпляр Echo для обработки HTTP-запросов.

	// Определение маршрутов
	e.GET("/tasks", handler.GetTasksHandler)
	e.POST("/tasks", handler.PostTaskHandler)
	e.PUT("/tasks/{id}", handler.UpdateHandler)
	e.PATCH("/tasks/{id}", handler.PatchHandler)
	e.DELETE("/tasks/{id}", handler.DeleteHandler)
	// Запуск сервера на порту 8080
	e.Logger.Fatal(e.Start(":8080")) // Запускаем сервер на порту 8080 и логируем ошибки, если они возникнут.
}
