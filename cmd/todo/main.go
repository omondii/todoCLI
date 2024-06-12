package main

import (
	"flag"
	"fmt"
	"os"
	//"/cmd/todo"
)

const (
	todoFile = ".todos.json"
)

func main() {
	// Define boolean flags with name(string), default value(bool), and usage description(string)
	add := flag.Bool("add", false, "add a new Todo")

	//parse the flags created
	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Sample ToDo")
	default:
		fmt.Fprintln(os.Stdout, "Invalid Command")
		os.Exit(0)
	}
}
