package middlewares

import (
	"log"
	"os"

	"github.com/KinitaL/go-crud/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Auth(c *fiber.Ctx) error {
	token := c.Cookies("token")
	db := connectToDB()
	var user models.User
	db.Where("access_token = ?", token).First(&user)
	if user.Access_token != "" {
		return c.Next()
	} else {
		return c.JSON("Wrong auth token")
	}
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
