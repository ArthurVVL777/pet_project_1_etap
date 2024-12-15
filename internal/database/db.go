package database

import (
	"gorm.io/driver/postgres" // Импортируем драйвер PostgreSQL для GORM.
	"gorm.io/gorm"            // Импортируем основной пакет GORM для работы с ORM.
	"log"                     // Импортируем пакет для логирования ошибок.
)

// DB - экспортируемая переменная для хранения подключения к базе данных.
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
