package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()

	//fmt.Println(db.ExistsTable("users"))
	//db.CreateTable(models.UserSchema, "users")

	//db.Ping()

	//db.TruncateTable("users")

	//user := models.CreateUser("Julian", "julian123", "julian7@gmail.com")
	//fmt.Println(user)

	//users := models.ListUsers()
	//fmt.Println(users)

	user := models.GetUser(1)
	fmt.Println(user)
	// user.UserName = "Pepe"
	// user.Save()

	user.Delete()

	fmt.Println(models.ListUsers())
	db.Close()
}
