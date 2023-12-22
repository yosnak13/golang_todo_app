package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {

	/*
		u := &models.User{}
		u.Name = "test"
		u.Email = "test@example.com"
		u.PassWord = "testtest"
		fmt.Println(u)

		u.CreateUser()
	*/

	fmt.Println(models.Db)
	u, _ := models.GetUser(1)
	fmt.Println(u)

	//u.Name = "test2"
	//u.Email = "test2@example.com"
	//u.UpdateUser()
	//u, _ = models.GetUser(1)
	//fmt.Println(u)

	u.DeleteUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)

}
