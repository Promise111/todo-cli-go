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
	fmt.Printf("%+v\n\n", todos)
}
