package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const FILENAME = "data.json"

func loadData() []todo {
	file, err := os.OpenFile(FILENAME, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file '%s'. Error: %v", FILENAME, err)
		os.Exit(1)
	}
	defer file.Close()

	stat, _ := file.Stat()
	if stat.Size() == 0 {
		defaultData := todo{Name: "My First TODO"}
		return []todo{defaultData}
	}

	var todos []todo
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&todos); err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding '%s'. ERROR: %v", FILENAME, err)
		os.Exit(1)
	}
	return todos

}

func saveData(todos []todo) error {
	file, err := os.Create(FILENAME)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(todos)
	if err != nil {
		return err
	}

	return nil
}
