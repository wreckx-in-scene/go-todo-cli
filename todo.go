package main

import "fmt"

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

var todos []Todo

func AddTodo(title string) {
	todo := Todo{
		ID:        len(todos) + 1,
		Title:     title,
		Completed: false,
	}
	todos = append(todos, todo)
}

func ListTodos() {
	for _, todo := range todos {
		fmt.Println(todo.ID, todo.Title, todo.Completed)
	}
}
