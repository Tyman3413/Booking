package main

import (
	"log"
	"os"

	"github.com/Tyman3413/Booking/database"
	"github.com/Tyman3413/Booking/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	// Подключение к базе данных
	database.Connect()

	// Подключение .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Определение порта сервера через .env файл
	port := os.Getenv("PORT")

	// Создание экземпляра приложения на Fiber v2
	app := fiber.New()

	// Установка рутов
	routes.Setup(app)

	// Запуск API сервера
	app.Listen(":" + port)
}
