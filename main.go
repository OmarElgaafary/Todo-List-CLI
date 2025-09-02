package main

import (
	"app/to-do/cli"
)

func main() {
	myToDo := initToDo()
	cli.ReadCLI(myToDo)
}