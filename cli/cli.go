package cli

import (
	"app/to-do/todo"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadCLI(list *todo.ToDoList) {
	args := os.Args[1:]

    if len(args) < 1 {
        fmt.Println("Usage: todo <command> [arguments]")
        os.Exit(1)
    }

    command := args[0]
	fmt.Println(args)
	switch command {
	case "get":
		
		if len(args) < 1 {
            fmt.Println("Usage: todo get <id|all>")
            os.Exit(1)
        }

		param := args[1]
		if param == "all" {
			fmt.Println(list)
			return
		}
		toDoById, err := todo.GetToDoById(list ,param)

		if err != nil {
			panic("Error fetching todo.")
		}
		fmt.Println(toDoById)
		return

	case "create":
		fmt.Println("Enter todo Id: ")
		inputID := takeCLIArgs()
		fmt.Println("Enter todo title: ")
		inputTitle := takeCLIArgs()
		fmt.Println("Enter todo content: ")
		inputContent := takeCLIArgs()

		_ , err := todo.New(*list, inputID, inputTitle, inputContent)

		if err != nil {
			panic("Error while creating todo.")
		}
	case "delete":

		if len(args) < 1 {
            fmt.Println("Usage: todo delete <id>")
            os.Exit(1)
        }

		param := args[1]
		err := todo.DeleteToDoById(*list ,param)

		if err != nil {
			panic("Error deleteing todo.")
		}
	case "mark":

		if len(args) != 3 {
            fmt.Println("Usage: todo mark <id> <check / uncheck>")
            os.Exit(1)
        }

		param := args[1]
		toDoById, err := todo.GetToDoById(list ,param)
		
		if err != nil {
			fmt.Println("Error retrieving todo.")
            os.Exit(1)
		}
		checkStatus := args[2]
		if checkStatus == "check" {
			toDoById.Done = true
			list.Save()
		}
    default:
        fmt.Println("Unknown command:", command)
        fmt.Println("Available commands: get")
        os.Exit(1)
	}


}

func takeCLIArgs() string {
	var userInput string

	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	fmtTxt := strings.TrimSuffix(userInput, "\n")
	fmtTxt = strings.TrimSuffix(fmtTxt, "\r")

	if userInput == "" {
		panic("Cannot use empty string as input.")
	}

	return fmtTxt
}