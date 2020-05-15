package domain

import "github.com/jinzhu/gorm"

type UserRepository interface {
	FindByID(id int) *User
	FindAll() []User
	Save(user User) (*User, error)
}

type User struct {
	gorm.Model
	FirstName  string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age string `json:"age"`
	Email string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Address string `json:"address"`
}
