package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Tyman3413/Booking/database"
	"github.com/Tyman3413/Booking/models"
	"github.com/Tyman3413/Booking/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Функция для валидации электронной почты
func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`) // Регулярное выражение
	return Re.MatchString(email)
}

// Обработчик для регистрации нового пользователя
func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User

	// Извлечение данных пользователя из тела запроса
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}

	// Если длина пароля меньше или равна 8 символам, функция возвращает ошибку
	if len(data["password"].(string)) <= 8 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be greater than 8 characters",
		})
	}

	// Проверка email, полученного из данных пользователя
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid email",
		})
	}

	// Проверка, существует ли пользователь с таким же email
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	// Создается новый объект User с данными пользователя (имя, email) и хешированным паролем
	user := models.User{
		Name:  data["name"].(string),
		Email: strings.TrimSpace(data["email"].(string)),
	}

	// Сохранение нового пользователя в базе данных.
	user.SetPassword(data["password"].(string))
	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}

	// В случае успешной регистрации функция возвращает успешный статус
	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Account was successfully created",
	})
}

// Обработчик для аутентификации пользователя
func Login(c *fiber.Ctx) error {
	var data map[string]string
	var user models.User

	// Извлечение данных пользователя (email и пароль) из тела запроса
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}

	// Выполняется запрос к базе данных, чтобы найти пользователя с указанным email
	database.DB.Where("email=?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "This email address doesn't exist. Do you want to create an account?",
		})
	}

	// Проверка пароля
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	// Если пароль верный, вызывается функция генерации JWT-токена
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	// Создается новый cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	// Функция возвращает успешный статус, устанавливает cookie в ответе и возвращает JSON-ответ
	return c.JSON(fiber.Map{
		"message": "Logined successfully",
		"user":    user,
	})
}

// Структура, используемая для хранения данных JWT-токена.
type Claims struct {
	jwt.StandardClaims
}
