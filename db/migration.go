package db

import (
	"github.com/f-katkit/simple-gin/forms"
)

func autoMigration() {
	db.AutoMigrate(&forms.User{})
}
