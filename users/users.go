package users

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/DdZ-Fred/fiber-server-1/dbUtils"
	"github.com/DdZ-Fred/fiber-server-1/models"
	"github.com/DdZ-Fred/fiber-server-1/password"
	"github.com/DdZ-Fred/fiber-server-1/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

func GetUsersSlice() []models.User {
	usersJsonBytes, _ := ioutil.ReadFile("users.json")
	var usersSlice []models.User
	json.Unmarshal(usersJsonBytes, &usersSlice)
	return usersSlice
}

func UsersRouter(app *fiber.App, db *gorm.DB) {
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

	// [JSON] GET ALL
	// users.Get("/", func(c *fiber.Ctx) error {
	// 	usersSlice := GetUsersSlice()
	// 	return c.JSON(usersSlice)
	// })

	// [DB] GET ALL
	users.Get("/", func(c *fiber.Ctx) error {
		var users []models.UserSafe

		pagination := dbUtils.GeneratePaginationFromRequest(c)
		offset := (pagination.Page - 1) * pagination.PerPage
		var totalRows int64 = 1

		db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Model(&models.User{}).Limit(pagination.PerPage).Offset(offset).Find(&users).Error; err != nil {
				// Returning any error rollbacks the transaction
				return err
			}
			if err := tx.Model(&models.User{}).Count(&totalRows).Error; err != nil {
				return err
			}

			// Returning nil commits the whole transaction
			return nil
		})

		dbUtils.SetPaginationTotalPages(&pagination, totalRows)

		return c.Status(200).JSON(fiber.Map{
			"data":       users,
			"pagination": pagination,
		})
	})

	// POST v2: user added to DB
	users.Post("/", func(c *fiber.Ctx) error {
		var payload PostPayload

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		validationErrors := PostValidate(payload)
		if validationErrors != nil {
			return validation.FailedValidationResponse(c, validationErrors)
		}

		birthDate, _ := time.Parse(time.RFC3339, payload.BirthDate)
		password, _ := password.EncryptPassword(payload.Password, &password.Params{
			Memory:      64 * 1024,
			Iterations:  3,
			Parallelism: 2,
			SaltLength:  16,
			KeyLength:   32,
		})

		newUser := models.User{
			Id:        uuid.New().String(),
			Fname:     payload.Fname,
			Lname:     payload.Lname,
			Email:     payload.Email,
			BirthDate: birthDate,
			Password:  password,
		}

		if err := db.Create(&newUser).Error; err != nil {
			// Read about error value: https://blog.golang.org/go1.13-errors#TOC_2.1.
			if pgErr, ok := err.(*pgconn.PgError); ok {
				switch pgErr.Code {
				case "23505":
					return c.Status(fiber.StatusConflict).JSON(fiber.Map{
						"originalError": pgErr,
						"code":          1001,
						"status":        "email_already_taken",
					})
				}
			}
			return c.Status(fiber.StatusInternalServerError).JSON(err)
			// return err
		}

		return c.Status(201).JSON(newUser)
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
