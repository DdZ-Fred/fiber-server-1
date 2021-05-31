package fiberserver1

import (
	"log"

	"github.com/DdZ-Fred/fiber-server-1/users"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()

	users.UsersRouter(app)

	log.Fatal(app.Listen(":3000"))
}
