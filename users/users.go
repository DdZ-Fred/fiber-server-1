package users

import (
	"encoding/json"
	"io/ioutil"

	"github.com/DdZ-Fred/fiber-server-1/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
		user, err := Find(usersSlice, func(user models.User, idx int) bool {
			return user.Id == c.Params("id")
		})

		if err != nil {
			return c.SendStatus(404)
		}

		return c.JSON(user)
	})

	users.Get("/", func(c *fiber.Ctx) error {
		usersSlice := GetUsersSlice()
		return c.JSON(usersSlice)
	})

	users.Post("/", func(c *fiber.Ctx) error {
		user, _ := GetParsedBody(c)

		user.Id = uuid.New().String()

		usersSlice := GetUsersSlice()
		usersSlice = append(usersSlice, user)

		usersSliceBytes, _ := json.Marshal(usersSlice)

		err := ioutil.WriteFile("users.json", usersSliceBytes, 0600)

		if err != nil {
			return err
		}
		return c.Status(201).JSON(user)
	})

	users.Put("/:id", func(c *fiber.Ctx) error {
		updatedUser, parsingError := GetParsedBody(c)

		if parsingError != nil {
			return parsingError
		}

		updatedUser.Id = c.Params("id")

		usersSlice := GetUsersSlice()

		userIdx, err := FindIndex(usersSlice, func(user models.User) bool {
			return user.Id == c.Params("id")
		})

		if err != nil {
			return c.SendStatus(404)
		}

		usersSlice[userIdx] = updatedUser
		usersSliceBytes, _ := json.Marshal(usersSlice)

		err2 := ioutil.WriteFile("users.json", usersSliceBytes, 0600)
		if err2 != nil {
			return err2
		}

		return c.JSON(updatedUser)
	})

	users.Delete("/:id", func(c *fiber.Ctx) error {
		usersSlice := GetUsersSlice()

		userIdx, err := FindIndex(usersSlice, func(user models.User) bool {
			return user.Id == c.Params("id")
		})

		if err != nil {
			return c.SendStatus(404)
		}

		usersSlice = RemoveIndex(usersSlice, userIdx)

		usersSliceBytes, _ := json.Marshal(usersSlice)

		err2 := ioutil.WriteFile("users.json", usersSliceBytes, 0600)

		if err2 != nil {
			return nil
		}

		return c.SendStatus(204)
	})

}
