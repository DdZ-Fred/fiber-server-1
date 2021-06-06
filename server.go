package fiberserver1

import (
	"log"

	"github.com/DdZ-Fred/fiber-server-1/db"
	"github.com/DdZ-Fred/fiber-server-1/users"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	db := db.InitDB()
	app := fiber.New()

	users.UsersRouter(app, db)

	log.Fatal(app.Listen(":3000"))
}
