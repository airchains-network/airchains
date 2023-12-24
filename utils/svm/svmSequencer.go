package svm

import (
	"airchains/utils"
	"fmt"
	"path/filepath"
)

func SolanaSequencerClone() {
	repoURL := "https://github.com/airchains-network/svm-sequencer-node.git"
	destinationPath := filepath.Join(".", "svm-sequencer-node")
	if err := utils.ChainRepoCloner(repoURL, destinationPath, "SVM-Sequencer"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
