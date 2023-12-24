package cosmwasm

import (
	"airchains/utils"
	"fmt"
	"path/filepath"
)

func CosmwasmSequencerClone() {
	repoURL := "https://github.com/airchains-network/cosmwasm-sequencer-node.git"
	destinationPath := filepath.Join(".", "cosmwasm-sequencer-node")
	if err := utils.ChainRepoCloner(repoURL, destinationPath, "SVM-Sequencer"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
