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

// Определяем структуру User для представления информации о пользователе
// type User struct {
// 	ID       int
// 	Name     string
// 	Email    string
// 	Password string
// }

func main() {
	database.Connect()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":" + port)
}
