package users

import (
	"fmt"
	"unicode"

	"github.com/DdZ-Fred/fiber-server-1/validation"
	"github.com/go-playground/validator/v10"
)

func validatePassword(fl validator.FieldLevel) bool {
	length := len(fl.Field().String())
	lower := 0
	upper := 0
	number := 0
	special := 0

	if length < 8 {
		return false
	}

	for _, c := range fl.Field().String() {
		switch {
		case unicode.IsNumber(c):
			number += 1
		case unicode.IsUpper(c):
			upper += 1
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special += 1
		case unicode.IsLetter(c) || c == ' ':
			lower += 1
		default:
			fmt.Println("Unkown character type: %s", c)
		}
	}
	return lower >= 1 && upper >= 1 && number >= 1 && special >= 1
}

type PostPayload struct {
	Fname     string `validate:"required,min=2,max=30" json:"fname"`
	Lname     string `validate:"required,min=2,max=30" json:"lname"`
	Email     string `validate:"required,email" json:"email"`
	Password  string `validate:"required,password" json:"password"`
	BirthDate string `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"birthDate"`
}

func PostValidate(payload PostPayload) []*validation.ValidationError {
	var errors []*validation.ValidationError
	validate := validator.New()
	validate.RegisterValidation("password", validatePassword)
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element validation.ValidationError
			element.FailedField = err.StructNamespace()
			element.ValidatorKey = err.Tag()
			element.ValidatorParam = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
