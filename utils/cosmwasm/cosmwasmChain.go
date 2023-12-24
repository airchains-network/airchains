package cosmwasm

import (
	"airchains/utils"
	"fmt"
	"path/filepath"
)

func CosmosChainClone() {
	repoURL := "https://github.com/airchains-network/rollup-cosmwasm.git"
	destinationFolder := "rollup-cosmwasm"
	destinationPath := filepath.Join(".", destinationFolder)
	if err := utils.ChainRepoCloner(repoURL, destinationPath, "SVM"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
