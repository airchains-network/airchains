package evm

import (
	"airchains/utils"
	"fmt"
	"path/filepath"
)

func EVMChainClone() {
	repoURL := "https://github.com/airchains-network/rollup-evm.git"
	destinationPath := filepath.Join(".", "rollup-evm")
	if err := utils.ChainRepoCloner(repoURL, destinationPath, "EVM"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
