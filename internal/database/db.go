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
	// Строка подключения к базе данных (DSN).
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error // Объявляем переменную для хранения ошибок.

	// Открываем подключение к базе данных с использованием GORM и драйвера PostgreSQL.
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil { // Проверяем, произошла ли ошибка при подключении к базе данных.
		log.Fatal("Failed to connect to database: ", err) // Если ошибка есть, логируем её и завершаем программу.
	}
}
