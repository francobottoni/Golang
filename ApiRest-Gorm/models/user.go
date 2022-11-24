package models

import (
	"gorm/db"
)

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

func MigrateUser() {
	db.Database.AutoMigrate(User{})
}
