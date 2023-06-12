package controller

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/Tyman3413/Booking/database"
	"github.com/Tyman3413/Booking/models"
	"github.com/Tyman3413/Booking/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateHotel(c *fiber.Ctx) error {
	var hotelpost models.Hotel
	if err := c.BodyParser(&hotelpost); err != nil {
		fmt.Println("Unable to parse body")
	}

	if err := database.DB.Create(&hotelpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Post is done",
	})
}

func AllPosts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit

	var total int64
	var gethotel []models.Hotel

	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&gethotel)
	database.DB.Model(&models.Hotel{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": gethotel,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	})
}

func DetailPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var hotelpost models.Hotel

	database.DB.Where("id=?", id).Preload("User").First(&hotelpost)

	return c.JSON(fiber.Map{
		"data": hotelpost,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	hotel := models.Hotel{
		Id: uint(id),
	}

	if err := c.BodyParser(&hotel); err != nil {
		fmt.Println("Unable to parse body")
	}

	database.DB.Model(&hotel).Updates(hotel)

	return c.JSON(fiber.Map{
		"message": "The post was successfully updated",
	})
}

func UniquePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := util.Parsejwt(cookie)

	var hotel []models.Hotel

	database.DB.Model(&hotel).Where("user_id=?", id).Preload("User").Find(&hotel)

	return c.JSON(hotel)
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	hotel := models.Hotel{
		Id: uint(id),
	}

	deleteQuery := database.DB.Delete(&hotel)

	if errors.Is(deleteQuery.Error, gorm.ErrRecordNotFound) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "The post was successfully deleted",
	})
}
