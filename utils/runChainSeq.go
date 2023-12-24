package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RunChainSeq(folderPath, scriptName string) error {
	// Get the absolute path of the specified folder
	absFolderPath, err := filepath.Abs(folderPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}

	// Check if the directory exists
	if _, err := os.Stat(absFolderPath); os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %v", absFolderPath)
	}

	// Change to the specified folder
	if err := os.Chdir(absFolderPath); err != nil {
		return fmt.Errorf("failed to change directory: %v", err)
	}

	pm2Args := []string{"start", scriptName}

	// Run the shell script
	cmd := exec.Command("pm2", pm2Args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running shell script %s in folder %s\n", scriptName, absFolderPath)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run shell script: %v", err)
	}

	err = os.Chdir("../")
	if err != nil {
		return fmt.Errorf("failed to change directory: %v", err)
	}

	fmt.Println("Shell script executed successfully")

	return nil
}
