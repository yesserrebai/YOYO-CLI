package outterStrcutureHelpers

import (
	"encoding/json"
	"fmt"
	"os"
)

func GenerateJsonFile(filePath string, jsonData interface{}) {
	// Create the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating JSON file: %v\n", err)
		return
	}
	defer file.Close()

	// Marshal the JSON data
	jsonBytes, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	// Write the JSON data to the file
	_, err = file.Write(jsonBytes)
	if err != nil {
		fmt.Printf("Error writing JSON data to file: %v\n", err)
		return
	}

	fmt.Printf("JSON file generated: %s\n", filePath)
}
