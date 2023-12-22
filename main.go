package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.PassWord = "testtest"
	fmt.Println(u)

	u.CreateUser()
}
