package middleware

import (
	"github.com/Tyman3413/Booking/util"
	"github.com/gofiber/fiber/v2"
)

// Функция проверки аутентификации пользователя
func IsAuthenticate(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	return c.Next()
}

/*
	Функция IsAuthenticate проверяет наличие и действительность токена аутентификации в куки запроса.
	В случае отсутствия или недействительности токена, возвращается ответ, указывающий на отсутствие
	аутентификации. В противном случае, запрос передается следующим обработчикам маршрута.
*/
