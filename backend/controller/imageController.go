package controller

import (
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

// Срез из всех букв английского алфавита для генерации имени изображения
var letters = []rune("abcdefghijklmnopqrstuvwxyz")

// Функция генерирует случайную строку из букв английского алфавита
func randLetter(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// Функция обработки и загрузки файлов
func Upload(c *fiber.Ctx) error {
	// Получение формы с файлами
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	// Получение списка файлов из формы
	files := form.File["image"]
	// Переменная fileName будет содержать имя файла после генерации
	fileName := ""

	// Сохранение файла на сервере
	for _, file := range files {
		fileName = randLetter(5) + "-" + file.Filename // Генерация случайного имени файла
		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return nil
		}
	}

	// Возвращение ответа в формате JSON. Ключ "url" содержит URL, по которому можно получить доступ к загруженному файлу.
	return c.JSON(fiber.Map{
		"url": "http://localhost:3000/api/uploads/" + fileName,
	})
}
