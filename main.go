package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

	// u := &models.User{}
	// u := new(models.User)
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(2)
	// fmt.Println(u)

	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(u.ID)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user, _ := models.GetUser(3)
	// user.CreateTodo("Third Todo")

	// todo, _ := models.GetTodo(1)
	// fmt.Println(todo)

	// todos, _ := models.GetTodos()
	// for _, todo := range todos {
	// 	fmt.Println(todo)
	// }

	// user, _ := models.GetUser(2)
	// todos, _ := user.GetTodosByUser()
	// for _, todo := range todos {
	// 	fmt.Println(todo)
	// }

	// todo, _ := models.GetTodo(1)
	// fmt.Println(todo)

	// todo.Content = "Update Todo"
	// todo.UpdateTodo()
	// updateTodo, _ := models.GetTodo(todo.ID)
	// fmt.Println(updateTodo)

	todo, _ := models.GetTodo(4)
	todo.DeleteTodo()
}
