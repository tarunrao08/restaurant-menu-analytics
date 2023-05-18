package main

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	// Create a log.txt file for testing
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	defer file.Close()

	// Write sample log entries
	_, err = file.WriteString("1 101\n")
	if err != nil {
		log.Fatalf("Failed to write log entry: %v", err)
	}
	_, err = file.WriteString("2 102\n")
	if err != nil {
		log.Fatalf("Failed to write log entry: %v", err)
	}
	_, err = file.WriteString("3 101\n")
	if err != nil {
		log.Fatalf("Failed to write log entry: %v", err)
	}
	_, err = file.WriteString("4 103\n")
	if err != nil {
		log.Fatalf("Failed to write log entry: %v", err)
	}
	_, err = file.WriteString("5 104\n")
	if err != nil {
		log.Fatalf("Failed to write log entry: %v", err)
	}
	_, err = file.WriteString("6 102\n")
	if err != nil {
		log.Fatalf("Failed to write log entry: %v", err)
	}
	_, err = file.WriteString("7 103\n")
	if err != nil {
		log.Fatalf("Failed to write log entry: %v", err)
	}
}

func teardown() {
	// Remove the log.txt file after testing
	err := os.Remove("log.txt")
	if err != nil {
		log.Printf("Failed to remove log file: %v", err)
	}
}

func TestMainFunction(t *testing.T) {
	// Capture stdout for testing output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the main function
	main()

	// Read captured stdout
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	expectedOutput := "Top 3 menu items consumed:\n1. FoodmenuID: 101, Count: 2\n2. FoodmenuID: 102, Count: 2\n3. FoodmenuID: 103, Count: 2\n"
	assert.Equal(t, expectedOutput, string(out))
}

func TestDuplicateMenuItems(t *testing.T) {
	// Modify the log.txt file to include duplicate menu items
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString("1 101\n")
	if err != nil {
		log.Fatalf("Failed to write log entry: %v", err)
	}
	_, err = file.WriteString("8 102\n")
	if err != nil {
		log.Fatalf("Failed to write log entry: %v", err)
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the main function
	main()

	// Read captured stdout
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	expectedError := "Error: Duplicate foodmenu IDs found\n"
	assert.Equal(t, expectedError, string(out))
}

func TestInvalidLogEntry(t *testing.T) {
	// Create a log.txt file for testing
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString("9\n")
	if err != nil {
		log.Fatalf("Failed to write log entry: %v", err)
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the main function
	main()

	// Read captured stdout
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	expectedError := "Invalid log entry: 9"
	assert.Equal(t, expectedError, string(out))
}
