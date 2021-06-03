package users

import (
	"errors"

	"github.com/DdZ-Fred/fiber-server-1/models"
	"github.com/gofiber/fiber/v2"
)

// Find {User} item in slice
func Find(slice []models.User, finder func(user models.User, idx int) bool) (models.User, error) {
	for index, value := range slice {
		if finder(value, index) {
			return value, nil
		}
	}
	return models.User{Id: "0", Fname: "", Lname: "", Email: ""}, errors.New("not-found")
}

// FindIndex {User} item in slice
func FindIndex(slice []models.User, finder func(user models.User) bool) (int, error) {
	for index, value := range slice {
		if finder(value) {
			return index, nil
		}
	}
	return 0, errors.New("not-found")
}

// GetParsedBody expecting a body of type {User}
func GetParsedBody(c *fiber.Ctx) (models.User, error) {
	var user models.User

	err := c.BodyParser(&user)

	if err != nil {
		return models.User{Id: "0", Fname: "", Lname: "", Email: ""}, err
	}

	return user, nil
}

// RemoveIndex: removes {User} at index `index`
func RemoveIndex(slice []models.User, index int) []models.User {
	return append(slice[:index], slice[index+1:]...)
}
