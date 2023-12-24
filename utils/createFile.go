// fileutil.go
package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// CreateFolderAndJSONFile creates a folder and a JSON file with the provided data.
func CreateFolderAndJSONFile(data map[string]interface{}, folderName, fileName string) {
	// Check if the folder exists; create it if not
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		if err := os.Mkdir(folderName, 0755); err != nil {
			fmt.Printf("Error creating folder: %v\n", err)
			os.Exit(1)
		}
	}

	// Combine folder and file paths
	jsonFilePath := filepath.Join(folderName, fileName)

	// Check if the JSON file exists
	if _, err := os.Stat(jsonFilePath); os.IsNotExist(err) {
		// Create a new JSON file
		createNewJSONFile(data, jsonFilePath)
	} else {
		// Update the existing JSON file
		updateExistingJSONFile(data, jsonFilePath)
	}

	fmt.Printf("Data has been saved into config.json: %s\n", jsonFilePath)
}

func createNewJSONFile(data map[string]interface{}, filename string) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile(filename, jsonData, 0644); err != nil {
		fmt.Printf("Error writing JSON file: %v\n", err)
		os.Exit(1)
	}
}

func updateExistingJSONFile(data map[string]interface{}, filename string) {
	// Read the existing JSON file
	existingData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading existing JSON file: %v\n", err)
		os.Exit(1)
	}

	// Unmarshal the existing JSON data
	var existingMap map[string]interface{}
	if err := json.Unmarshal(existingData, &existingMap); err != nil {
		fmt.Printf("Error decoding existing JSON data: %v\n", err)
		os.Exit(1)
	}

	// Update the existing map with new data
	for key, value := range data {
		existingMap[key] = value
	}

	// Encode the updated data and write it back to the file
	updatedData, err := json.MarshalIndent(existingMap, "", "  ")
	if err != nil {
		fmt.Printf("Error encoding updated JSON: %v\n", err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile(filename, updatedData, 0644); err != nil {
		fmt.Printf("Error writing updated JSON file: %v\n", err)
		os.Exit(1)
	}
}

func CreateENV(data string, folderName, fileName string) {

	jsonFilePath := filepath.Join(folderName, fileName)

	fmt.Printf("Data has been saved into .env: %s\n", jsonFilePath)
}

func ENVSetup(filePath string, exePort string, daTypeSend string) {
	ExecutionClientRPC := fmt.Sprintf("http://localhost:%s", exePort)
	ExecutionClientTRPC := "http://localhost:1317"
	ExecutionClientJsonRPC := "http://localhost:26657"
	DaClientRPC := fmt.Sprintf("http://localhost:5050/%s", daTypeSend)
	SettlementClientRPC := "http://localhost:8080"

	envFileData := []byte(fmt.Sprintf("ExecutionClientRPC=%s\nExecutionClientTRPC=%s\nExecutionClientJsonRPC=%s\nDaClientRPC=%s\nSettlementClientRPC=%s\n", ExecutionClientRPC, ExecutionClientTRPC, ExecutionClientJsonRPC, DaClientRPC, SettlementClientRPC))
	filePathRoot := fmt.Sprintf("%s/.env", filePath)
	envWriteErr := os.WriteFile(filePathRoot, envFileData, 0666)
	if envWriteErr != nil {
		fmt.Println("Error writing to .env file:", envWriteErr)
	}
}
