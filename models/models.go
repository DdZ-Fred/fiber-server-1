package models

import "time"

type User struct {
	Id        string    `gorm:"primaryKey" json:"id,omitempty"`
	Fname     string    `json:"fname"`
	Lname     string    `json:"lname"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password,omitempty"`
	BirthDate time.Time `json:"birthDate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
