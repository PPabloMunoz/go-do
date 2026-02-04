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
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(todos); err != nil {
		return fmt.Errorf("failed to encode todos: %w", err)
	}

	count := len(todos)
	countStr := fmt.Sprintf("%d", count)
	if count == 1 {
		fmt.Printf("  %s  Saved %s todo to %s\n", checkmarkStyle.Render("✓"), greenStyle.Render(countStr), FILENAME)
	} else {
		fmt.Printf("  %s  Saved %s todos to %s\n", checkmarkStyle.Render("✓"), greenStyle.Render(countStr), FILENAME)
	}
	return nil
}
