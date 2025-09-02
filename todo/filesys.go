package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteToFile(filePath string, data any) error {
	dataJSON, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	os.WriteFile(filePath, dataJSON, 0644)
	return nil
}

func ParseJSONfromFile(filePath string) (ToDoList, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error while parsing JSON data.")
		return []Todo{}, err
	}

	var storeJSON ToDoList
	json.Unmarshal(data, &storeJSON)

	return storeJSON, nil
}