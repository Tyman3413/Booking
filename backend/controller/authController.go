package controller

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/Tyman3413/Booking/database"
	"github.com/Tyman3413/Booking/models"
	"github.com/gofiber/fiber/v2"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}

	if len(data["password"].(string)) <= 8 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be greater than 8 characters",
		})
	}

	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid email",
		})
	}

	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)

	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already exists",
		})
	}

	user := models.User{
		Name:  data["name"].(string),
		Email: strings.TrimSpace(data["email"].(string)),
	}

	user.SetPassword(data["password"].(string))
	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Account was successfully created",
	})
}
