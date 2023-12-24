package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func ChainSetup(folderPath string) {
	// Specify the folder path and script name
	scriptName := "setup.sh"
	if err := os.Chdir(folderPath); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Run the shell script
	cmd := exec.Command("sh", scriptName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running shell script %s in folder %s\n", scriptName, folderPath)

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err := os.Chdir("../")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Chain Built Successfully")
}
