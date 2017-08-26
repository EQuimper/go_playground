package main

import (
	"fmt"
)

var currentID int

var todos Todos

func init() {
	RepoCreateTodo(Todo{Name: "First todo"})
	RepoCreateTodo(Todo{Name: "Second todo"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}

	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentID++
	t.Id = currentID
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
