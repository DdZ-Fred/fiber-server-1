package models

import "time"

type UserSafe struct {
	Id        string    `gorm:"primaryKey" json:"id,omitempty"`
	Fname     string    `json:"fname"`
	Lname     string    `json:"lname"`
	Email     string    `gorm:"unique" json:"email"`
	BirthDate time.Time `json:"birthDate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	Id        string    `gorm:"primaryKey" json:"id,omitempty"`
	Fname     string    `json:"fname"`
	Lname     string    `json:"lname"`
	Email     string    `gorm:"unique" json:"email"`
	BirthDate time.Time `json:"birthDate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Password  string    `json:"password,omitempty"`
}
