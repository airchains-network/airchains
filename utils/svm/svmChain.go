package svm

import (
	"airchains/utils"
	"fmt"
	"path/filepath"
)

func SolanaChainClone() {
	repoURL := "https://github.com/airchains-network/rollup-svm.git"
	destinationFolder := "rollup-svm"
	destinationPath := filepath.Join(".", destinationFolder)
	if err := utils.ChainRepoCloner(repoURL, destinationPath, "SVM"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	utils.ChainSetup(destinationFolder)
}
