package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"pet_project_1_etap/internal/userService"
)

var DB *gorm.DB

// InitDB инициализирует базу данных.
func InitDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"

	var err error // Объявляем переменную для хранения ошибок.

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err) // Логируем ошибку и завершаем программу.
	}
}
func Migrate() {
	if err := DB.AutoMigrate(&userService.User{}); err != nil {
		log.Fatalf("Automigrate failed: %v", err)
	}
}
