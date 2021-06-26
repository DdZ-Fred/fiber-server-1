package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type UserSafe struct {
	Id        string    `gorm:"primaryKey" json:"id,omitempty"`
	Fname     string    `gorm:"notNull" json:"fname"`
	Lname     string    `gorm:"notNull" json:"lname"`
	Email     string    `gorm:"notNull;unique" json:"email"`
	BirthDate time.Time `gorm:"notNull" json:"birthDate"`
	CreatedAt time.Time `gorm:"notNull" json:"createdAt"`
	UpdatedAt time.Time `gorm:"notNull" json:"updatedAt"`

	ConfirmedAt null.Time `gorm:"default:null" json:"confirmedAt"`
	ActivatedAt null.Time `gorm:"default:null" json:"activatedAt"`
	Activated   bool      `gorm:"notNull;default:false" json:"activated"`
}

type User struct {
	Id        string    `gorm:"primaryKey" json:"id,omitempty"`
	Fname     string    `gorm:"notNull" json:"fname"`
	Lname     string    `gorm:"notNull" json:"lname"`
	Email     string    `gorm:"notNull;unique" json:"email"`
	BirthDate time.Time `gorm:"notNull" json:"birthDate"`
	CreatedAt time.Time `gorm:"notNull" json:"createdAt"`
	UpdatedAt time.Time `gorm:"notNull" json:"updatedAt"`
	Password  string    `gorm:"notNull" json:"password,omitempty"`

	ConfirmedAt null.Time `gorm:"default:null" json:"confirmedAt"`
	ActivatedAt null.Time `gorm:"default:null" json:"activatedAt"`
	Activated   bool      `gorm:"notNull;default:false" json:"activated"`
}
