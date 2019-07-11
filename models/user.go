package models

import (
	"github.com/f-katkit/simple-go-gin-boilerplate/db"
	"github.com/f-katkit/simple-go-gin-boilerplate/forms"
)

type User forms.User

func (s User) GetAll() ([]User, error) {
	db := db.GetDB()
	var u []User

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
