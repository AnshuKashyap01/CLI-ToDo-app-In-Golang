package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
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

func(todos *Todos) toggle(index int) error{
	t := *todos

	if err:= t.validIndex(index) ; err!=nil { //Shortcut to write Intialization and IF condition 
		return err;
	}

	isCompleted := t[index].Completed //true or false

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime 
	}

	t[index].Completed=!isCompleted
	 
	return nil;

}

func(todos *Todos ) edit(index int , title string) error{
	t:= *todos
	if err := t.validIndex(index) ; err!=nil{
		return err
	}
	t[index].Title = title

	return nil
}

func(todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#","Title","Completed","Created At","Completed At")
	for index,t := range *todos {
		completed := "❌"
		completedAt := ""
		if t.Completed {
			completed = "✅"
			if t.CompletedAt!= nil{
				completedAt= t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index),t.Title,completed,t.CreatedAt.Format(time.RFC1123),completedAt)
	}

	table.Render()
}





