package property

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Property struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

func InitialMigration() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
	var DB_NAME = os.Getenv("DB_NAME")
	var DB_USER = os.Getenv("DB_USER")
	var DB_USER_PASSWORD = os.Getenv("DB_USER_PASSWORD")
	var MYSQL_NETWORK_PORT = os.Getenv("MYSQL_LOCAL_PORT")

	var DNS = fmt.Sprintf("%s:%s@tcp(mysql_server:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_USER_PASSWORD, MYSQL_NETWORK_PORT, DB_NAME)
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	DB.AutoMigrate(&Property{})
}

func CreateProperty(c *fiber.Ctx) error {
	property := new(Property)
	if err := c.BodyParser(property); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Create(&property)
	return c.JSON(&property)
}

func GetProperty(c *fiber.Ctx) error {
	id := c.Params("id")
	var property Property
	DB.Find(&property, id)
	return c.JSON(&property)
}

func GetProperties(c *fiber.Ctx) error {
	var properties []Property
	DB.Find(&properties)
	return c.JSON(&properties)
}

func DeleteProperty(c *fiber.Ctx) error {
	id := c.Params("id")
	var property Property
	DB.First(&property, id)
	DB.Delete(&property)
	return c.JSON(&property)
}
