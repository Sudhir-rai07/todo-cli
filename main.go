package main

import (
	"flag"
	"log"

	"github.com/Sudhir-rai07/todo-cli/data"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("a", false, "add a new todo")
	todo := flag.String("t", "", "todo text")
	complete := flag.Int("c", 0, "mark a todo as completed")
	del := flag.Int("d", 0, "delete a todo")
	list := flag.Bool("ls", false, "list all todos")

	flag.Parse()

	todos := &data.Todos{}

	err := todos.Load(todoFile)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	switch {
	case *add:
		todos.Add(*todo)
		err := todos.Store(todoFile)
		if err != nil {
			log.Fatal(err.Error())
		}
	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil {
			log.Fatalf("Faild to complete, %v", err)
		}
		err = todos.Store(todoFile)
		if err != nil {
			log.Fatalf("Faild to Store, %v", err)
		}

	case *del > 0:
		err := todos.Delete(*del)
		if err != nil {
			log.Fatalf("Faild to Delte, %v", err)
		}
		err = todos.Store(todoFile)
		if err != nil {
			log.Fatalf("Faild to Store{deletion}, %v", err)
		}
	case *list:
		todos.PrintTodos()
	}

}
