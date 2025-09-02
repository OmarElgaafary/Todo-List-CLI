package todo

import (
	"errors"
	"fmt"
	"slices"
	"time"
)

type Todo struct {
	Id			string
	Title     	string
	Content   	string
	Done        bool
	CreatedAt 	string
}


var filePath string = "my-to-do-list.json"


type ToDoList []Todo

func (t *ToDoList) Save() error {
	err := WriteToFile("my-to-do-list.json", *t)
	if err != nil{ 
		fmt.Println("Error while writing to file")
		return err
	}

	fmt.Printf("Updated file '%v' successfully!\n", "my-to-do-list.json")
	return nil
}


func New(toDoArray ToDoList ,toDoID string, toDoTitle string, toDoContent string) (*Todo, error) {
	newToDo := &Todo{
		Id: toDoID,
		Title: toDoTitle,
		Content: toDoContent,
		Done: false,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	toDoArray = append(toDoArray, *newToDo)
	err := toDoArray.Save()

	if err != nil {
		return &Todo{"", "", "", false, ""}, err
	}

	fmt.Println("Created and added to do successfully!")
	return newToDo, nil
}

func MarkAsDone(toDoArray *ToDoList ,t *Todo) {
	defer toDoArray.Save()
	if (!t.Done) {
		t.Done = true
		fmt.Println("The task was marked as done!", t, toDoArray)
		return
	}
	fmt.Println("Task already marked as done.")
}

func UnMarkAsDone(toDoArray *ToDoList ,t *Todo) {
	defer toDoArray.Save()
	if (t.Done) {
		t.Done = false
		fmt.Println("The task was unmarked.", t, toDoArray)
		return
	}
	fmt.Println("Task already unmarked.")
}


func AddToDo(toDoArray ToDoList ,todo *Todo) error {
	if (!slices.Contains(toDoArray, *todo)) {
		toDoArray = append(toDoArray, *todo)
		fmt.Println(toDoArray)
		return nil
	}
	
	return errors.New("Slice already in 'toDoArray'")
}

func DeleteToDoById(toDoArray ToDoList ,todoId string) error {
	todo, err := GetToDoById(&toDoArray, todoId)

	if err != nil {
		return err
	}

	if (!slices.Contains(toDoArray, *todo)) {
		return errors.New("Todo not found in 'toDoArray'")
	}

	toDoIndex := slices.Index(toDoArray, *todo)
	toDoArray = slices.Delete(toDoArray, toDoIndex, toDoIndex + 1)
	
	fmt.Printf("My todo list deleted successfully.\n%v", toDoArray)
	toDoArray.Save()
	return nil
}

func GetToDoById(toDoArray *ToDoList ,todoId string) (*Todo, error) {
	for i := range *toDoArray {
		if ((*toDoArray)[i].Id == todoId) {
			return &(*toDoArray)[i], nil
		}
	}

	return nil, errors.New("Todo not found")
}


