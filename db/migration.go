package db

import (
	"github.com/f-katkit/simple-go-gin-boilerplate/forms"
)

func autoMigration() {
	db.AutoMigrate(&forms.User{})
}
