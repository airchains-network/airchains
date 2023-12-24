package evm

import (
	"fmt"
	"os"
	"os/exec"
)

func runMultiNodeEVMShellScript(folderPath, scriptName string) error {
	// Change to the specified folder
	if err := os.Chdir(folderPath); err != nil {
		return fmt.Errorf("failed to change directory: %v", err)
	}

	// Run the shell script
	cmd := exec.Command("sh", scriptName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running shell script %s in folder %s\n", scriptName, folderPath)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run shell script: %v", err)
	}

	fmt.Println("Shell script executed successfully")

	return nil
}

func RunMultiNodeEVMChain() {
	// Specify the folder path and script name
	folderPath := "evm-chain"
	scriptName := "multi-node.sh"

	// Call the function to run the shell script
	if err := runMultiNodeEVMShellScript(folderPath, scriptName); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
