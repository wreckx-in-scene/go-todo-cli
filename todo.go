package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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
	if len(todos) == 0 {
		fmt.Println("No todos yet!")
		return
	}

	for _, todo := range todos {
		status := "[ ]"
		if todo.Completed {
			status = "[X]"
		}
		fmt.Println(todo.ID, status, todo.Title)
	}
}

func SaveTodos() {
	data, err := json.Marshal(todos)
	if err != nil {
		fmt.Println("Error saving todos: ", err)
		return
	}

	os.WriteFile("todos.json", data, 0644)
}

// function to load todos
func LoadTodos() {
	data, err := os.ReadFile("todos.json")
	if err != nil {
		return
	}

	json.Unmarshal(data, &todos)
}

// function to mark todo as done
func MarkTodo(id int) {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = true
			fmt.Println("Marked as done!")
			return
		}
	}

	fmt.Println("Todo not found")
}

func DeletTodo(id int) {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("Todo deleted!")
			return
		}
	}

	fmt.Println("Todo not found")
}
