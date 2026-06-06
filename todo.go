package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title       string     "json:title"
	Completed   bool       "json:completed"
	CreatedAt   time.Time  "json:createdAt"
	CompletedAt *time.Time "json:completedAt"
}

type Todos []Todo

func (t *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*t = append(*t, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		completedAt := time.Now()
		t[index].Completed = true
		t[index].CompletedAt = &completedAt
	} else {
		t[index].CompletedAt = nil
		t[index].Completed = false
	}
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	if title == "" {
		return errors.New("Error: title is require")
	}

	t[index].Title = title
	return nil
}
