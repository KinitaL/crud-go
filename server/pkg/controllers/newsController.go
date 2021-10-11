package controllers

import (
	"log"
	"os"

	"github.com/KinitaL/go-crud/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Get(c *fiber.Ctx) error {
	db := connectToDB()
	db.AutoMigrate(&models.News{})
	var news []models.News
	db.Find(&news)
	return c.JSON(news)
}

func Post(c *fiber.Ctx) error {
	db := connectToDB()
	var news models.News
	c.BodyParser(&news)
	db.Create(&news)
	return c.JSON(news)
}

func Put(c *fiber.Ctx) error {
	db := connectToDB()
	var news models.News
	c.BodyParser(&news)
	db.Model(models.News{}).Where("id=?", c.Params("id")).Updates(&news)
	return c.JSON(news)
}

func Delete(c *fiber.Ctx) error {
	db := connectToDB()
	db.Model(models.News{}).Where("id = ?", c.Params("id")).Delete(c.Params("id"))
	return c.JSON("Deleted")
}

func envLoad() {
	if err := godotenv.Load("./db.env"); err != nil {
		log.Fatal("No .env file found")
	}
}

func envManage(key string) string {

	value, exists := os.LookupEnv(key)
	if exists {
		return value
	} else {
		log.Fatal("No such env")
		return ""
	}
}

func connectToDB() *gorm.DB {
	envLoad()
	host := envManage("POSTGRES_HOST")
	port := envManage("POSTGRES_PORT")
	user := envManage("POSTGRES_USER")
	pass := envManage("POSTGRES_PASSWORD")
	dbname := envManage("POSTGRES_DB")
	helper := "host=" + host + " port=" + port + " user=" + user + " password=" + pass + " dbname=" + dbname + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(helper), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
