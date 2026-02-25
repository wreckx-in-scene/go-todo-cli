package main

import (
	"fmt"
	"os"
)

func main() {
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
	case "list":
		ListTodos()
	default:
		fmt.Println("Unknown command")
	}
}
