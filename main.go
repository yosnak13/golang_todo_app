package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	/*
		u := &models.User{}
		u.Name = "test"
		u.Email = "test@example.com"
		u.PassWord = "testtest"
		fmt.Println(u)

		u.CreateUser()
	*/

	/*
		fmt.Println(models.Db)
		u, _ := models.GetUser(1)
		fmt.Println(u)

		u.Name = "test2"
		u.Email = "test2@example.com"
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

		u.DeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/

	/*
		user, _ := models.GetUser(2)
		user.CreateTodo("FirstTodo")
	*/

	/*
		t, _ := models.GetTodo(1)
		fmt.Println(t)
	*/

	/*
		user, _ := models.GetUser(2)
		user.CreateTodo("SecondTodo")

		todos, _ := user.ListTodosByUser()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	//todos, _ := models.GetTodo()
	//for _, v := range todos {
	//	fmt.Println(v)
	//}

	t, _ := models.GetTodo(1)
	fmt.Println("-------------------")
	fmt.Println(t)
	t.Content = "Update Todo"
	t.UpdateTodo()

	updatedTodo, _ := models.GetTodo(1)
	fmt.Println("-------------------")
	fmt.Println(updatedTodo)

}
