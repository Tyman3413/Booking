package database

import (
	"log"
	"os"

	"github.com/Tyman3413/Booking/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Объявление глобальной переменной DB
var DB *gorm.DB

// Функция установки соединения с базой данных и настройки моделей
func Connect() {
	// Закрузка переменных из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Получение строки подключения к базе данных PostgreSQL
	dsn := os.Getenv("DSN")

	// Установка соединения с базой данных PostgreSQL с помощью функции Open
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	} else {
		log.Println("Successful connection to the database")
	}

	// Присвоение открытого соединения базы данных
	DB = database

	// Автомиграция
	database.AutoMigrate(
		&models.User{},
		&models.Hotel{},
	)

	/*
		Метод AutoMigrate принимает список моделей и создает соответствующие
		таблицы в базе данных, если они еще не существуют.
	*/
}
