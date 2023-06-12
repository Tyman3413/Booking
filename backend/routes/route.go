package routes

import (
	"github.com/Tyman3413/Booking/controller"
	"github.com/Tyman3413/Booking/middleware"
	"github.com/gofiber/fiber/v2"
)

// Функция настройки маршрутов и обработчиков для веб-приложения
func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register) // Регистрация пользователя
	app.Post("/api/login", controller.Login)       // Аутентификация пользователя

	app.Use(middleware.IsAuthenticate) // Проверка аутентификации пользователя перед обработкой других запросов

	app.Post("/api/post", controller.CreateHotel)            // Создание карточки Отеля
	app.Get("/api/allpost", controller.AllPosts)             // Выборка всех объектов сущности Отель
	app.Get("/api/allpost/:id", controller.DetailPost)       // Выборка объекта сущности Отель под определенным id
	app.Put("/api/updatepost/:id", controller.UpdatePost)    // Обновление информации об объекте сущности Отель под определенным id
	app.Get("/api/uniquepost", controller.UniquePost)        // Выборка всех объектов сущности Отель залогиненного пользователя
	app.Delete("/api/deletepost/:id", controller.DeletePost) // Удаление объекта сущности Отель под определенным id
	app.Post("/api/upload-image", controller.Upload)         // Загрузка изображения на сервер
	app.Static("/api/uploads", "./uploads")                  // Предоставление статичных файлов из дирректории ./uploads
}
