package main

import (
	"app/to-do/todo"
	"os"
)

func initToDo() *todo.ToDoList {
	filePath := "my-to-do-list.json"

	var myToDoList todo.ToDoList = []todo.Todo{}

	_ , err := os.Stat("my-to-do-list.json")

	if err == nil {
		list, err := todo.ParseJSONfromFile(filePath)

		if err != nil {
			panic(err)
		}

		myToDoList = list
	}

	return &myToDoList
}