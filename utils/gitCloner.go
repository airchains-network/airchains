package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func ChainRepoCloner(repoURL, destinationPath string, ChainType string) error {
	fmt.Printf("Cloning %s Chain ⏱️\n", ChainType)

	if err := os.RemoveAll(destinationPath); err != nil {
		return fmt.Errorf("error removing existing destination directory: %v", err)
	}

	// Use "git clone" command to clone the GitHub repository
	cmd := exec.Command("git", "clone", repoURL, destinationPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Cloning GitHub repository: %s to %s\n", repoURL, destinationPath)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to clone repository: %v", err)
	}

	fmt.Println("Repository cloned successfully")

	return nil

}
