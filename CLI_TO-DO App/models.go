package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo //we are creating slice of Todo data structure to store multiple task in Todos.

// Methods--->
func (todos *Todos) add(title string) { //we are defining a method for Todos slice to add a Task.
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo) //appending the Task in Todos.

}
func (todos *Todos) validIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos
	err := t.validIndex(index)
	if err != nil {
		return err
	}

	*todos = append(t[:index],t[index+1:]...)

	return nil;

}





