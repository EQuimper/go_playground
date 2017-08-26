package main

import (
	"time"
)

// Todo a todo
type Todo struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos list of todos
type Todos []Todo
