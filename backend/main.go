package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

func main() {
	// Подключение к базе данных PostgreSQL
	connection := "user=postgres password=admin dbname=booking sslmode=disable"

	// Открываем соединение с базой данных
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверка соединения с базой данных
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successful connection to the database")

	app := fiber.New()

	app.Get("/page", func(c *fiber.Ctx) error {
		return c.SendString("Working")
	})

	log.Fatal(app.Listen(":4000"))
}
