package utils

import (
	"fmt"
	"path/filepath"
)

func DaRPC() {
	repoURL := "https://github.com/airchains-network/da-client"
	destinationFolder := "da-rpc"
	destinationPath := filepath.Join(".", destinationFolder)
	if err := ChainRepoCloner(repoURL, destinationPath, "SVM"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
