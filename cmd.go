package main

import (
	"flag"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "Add", "", "Add a new todo, specify title")
	flag.StringVar(&cf.Edit, "Edit", "", "Edit a todo, specify index and title. id:new_title")
	flag.IntVar(&cf.Del, "Del", -1, "Specify a todo by index to delete")
	flag.IntVar(&cf.Toggle, "Toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "List", false, "List all todos")

	flag.Parse()

	return &cf
}
