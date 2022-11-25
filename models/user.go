package models

type User struct {
	ID    uint `json:"id" gorm:"primary_key"`
	Uuid  string
	Name  string `json:"name"`
	Email string `json:"email"`
}
