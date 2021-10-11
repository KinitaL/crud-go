package controllers

import (
	"time"

	"github.com/KinitaL/go-crud/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	db := connectToDB()
	db.AutoMigrate(&models.User{})
	var user models.User
	c.BodyParser(&user)
	db.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	db := connectToDB()

	//data from user
	var user models.User
	c.BodyParser(&user)

	//data from DB
	var dbuser models.User
	db.Where("access_token = ?", user.Access_token).First(&dbuser)

	//compare
	if user.Access_token == dbuser.Access_token {
		cookie := new(fiber.Cookie)
		cookie.Name = "token"
		cookie.Value = user.Access_token
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.Cookie(cookie)
		return c.JSON("You're logged in")
	} else {
		return c.JSON("wrong auth data")
	}
}

func DeleteUser(c *fiber.Ctx) error {
	db := connectToDB()
	var user models.User
	c.BodyParser(&user)
	db.Where("access_token = ?", &user.Access_token).Delete(&user)
	return c.JSON("user was deleted")
}
