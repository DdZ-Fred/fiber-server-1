package utils

import (
	"errors"

	"github.com/DdZ-Fred/fiber-server-1/models"
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
