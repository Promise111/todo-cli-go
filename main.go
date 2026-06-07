package main

import (
	// "fmt"
)

func main() {
	var todos = Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	commandFlags := NewCmdFlags()
	commandFlags.Execute(&todos)
	// todos.add("Build a todo cli")
	// todos.add("Boil water")
	// // todos.add("Make oats")
	// fmt.Printf("%+v\n\n", todos)
	// todos.delete(0)
	// todos.toggle(1)
	// todos.edit(0, "Bake bread")
	// todos.list()
	storage.Save(todos)
	// fmt.Printf("%+v\n\n", todos)
}
