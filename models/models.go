package models

type User struct {
	Id    string `json:"id,omitempty"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Email string `json:"email"`
}
