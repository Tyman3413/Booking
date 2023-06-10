package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

// Определяем структуру User для представления информации о пользователе
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

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
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	fmt.Println("Successful connection to the database!")

	// Создаем экземпляр сервера Fiber
	app := fiber.New()

	// Создаем обработчик POST-запроса для регистрации пользователя
	app.Post("/register", func(c *fiber.Ctx) error {
		// Получаем данные пользователя из запроса
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
		}

		// Добавляем нового пользователя в базу данных
		result, err := db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to register user"})
		}

		// Получаем ID нового пользователя
		userID, _ := result.LastInsertId()

		// Возвращаем ответ об успешной регистрации
		return c.JSON(fiber.Map{"message": "User registered successfully", "userID": userID})
	})

	log.Fatal(app.Listen(":4000"))
}
