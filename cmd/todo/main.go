package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/omondii/todoCLI"
)

const (
	todoFile = ".todos.json"
)

/* main prog entry point
 */
func main() {
	// Define boolean flags with name(string), default value(bool),
	// and usage description(string)
	add := flag.Bool("add", false, "Add new todo")
	complete := flag.Int("complete", 0, "Mark a todo as completed")
	del := flag.Int("del", 0, "Delete a todo")
	list := flag.Bool("list", false, "List all todos")

	flag.Parse()

	todos := &todo.Todos{}
	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error)
		os.Exit(1)
	}

	/*  */
	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		todos.Add(task)
		err = todos.Save(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Save(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *del > 0:
		err := todos.Delete(*del)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *list:
		todos.Print()
	default:
		fmt.Fprintln(os.Stdout, "Invalid Command")
		os.Exit(0)
	}
}

/* getInput: Captures command line arguments entered by user
 */
func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	if len(scanner.Text()) == 0 {
		return "", errors.New("Can not add an empty string")
	}
	return scanner.Text(), nil
}
