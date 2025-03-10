package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var dataUser = []User{} // arr เปล่า struct User

// struct json ที่รีบจากหน้าบ้าน
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

func main() {
	app := fiber.New()

	//allow port
	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173/",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// ฟังชั่น รับ random user top 10
	app.Post("/", func(c *fiber.Ctx) error {
		p := new([]User)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		dataUser = *p
		return c.SendString("random success")
	})

	// ฟังชั่น get random user top 10
	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(dataUser)
	})

	// ฟังชั่น get random user top 10
	app.Delete("/", func(c *fiber.Ctx) error {
		dataUser = []User{}
		return c.SendString("random clear")
	})

	app.Listen(":3000")
}
