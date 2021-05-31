package users

import (
	"encoding/json"
	"io/ioutil"

	"github.com/DdZ-Fred/fiber-server-1/models"
	"github.com/DdZ-Fred/fiber-server-1/utils"
	"github.com/gofiber/fiber/v2"
)

func GetUsersSlice() []models.User {
	usersJsonBytes, _ := ioutil.ReadFile("users.json")
	var usersSlice []models.User
	json.Unmarshal(usersJsonBytes, &usersSlice)
	return usersSlice
}

func UsersRouter(app *fiber.App) {
	users := app.Group("/users")

	users.Get("/:id", func(c *fiber.Ctx) error {
		usersSlice := GetUsersSlice()
		user, err := utils.Find(usersSlice, func(user models.User, idx int) bool {
			return user.Id == c.Params("id")
		})

		if err != nil {
			return c.SendStatus(404)
		}

		return c.JSON(user)
	})
	// users.Get("/", func(c *fiber.Ctx) error {

	// })
	// users.Post("/", func(c *fiber.Ctx) error {

	// })

}
