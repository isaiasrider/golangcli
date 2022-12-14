package main

import (
	"fmt"
	"os"
	"strings"
	"todo"
)

const todoFileName = "todo.json"

func main() {

	l := &todo.List{}
	// executa o get se nada der errado
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Println("fudeu")
		os.Exit(1)
	}
	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
		fmt.Println(l)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

}
