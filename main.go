package main

import (
	"fmt"
)

func main() {
	var todos = Todos{}
	todos.add("Boil water")
	todos.add("Make oats")
	fmt.Printf("%+v\n\n", todos)
	// todos.delete(0)
	todos.toggle(1)
	todos.edit(0, "Bake bread")
	fmt.Printf("%+v\n\n", todos)
}
