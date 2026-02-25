package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//load todos into the todo slice
	LoadTodos()

	if len(os.Args) < 2 {
		fmt.Println("Welcome to Todo CLI")
		fmt.Println("Commands:")
		fmt.Println("add 'task'	- Add a new todo")
		fmt.Println("list - Show all todos")
		fmt.Println("done 1 - Mark todo as complete")
		fmt.Println("delete 1 - Delete a todo")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		AddTodo(os.Args[2])
		SaveTodos()
		fmt.Println("Todo added!")
	case "list":
		ListTodos()
	case "done":
		id, _ := strconv.Atoi(os.Args[2])
		MarkTodo(id)
		SaveTodos()
	case "delete":
		id, _ := strconv.Atoi(os.Args[2])
		DeletTodo(id)
		SaveTodos()
	default:
		fmt.Println("Unknown command")
	}
}
