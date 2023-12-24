package evm

import (
	"airchains/utils"
	"fmt"
	"path/filepath"
)

func EvmSequencerClone() {
	repoURL := "https://github.com/airchains-network/evm-sequencer-node.git"
	destinationPath := filepath.Join(".", "evm-sequencer-node")
	if err := utils.ChainRepoCloner(repoURL, destinationPath, "EVM-Sequencer"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
