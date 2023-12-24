package utils

import (
	"fmt"
	"path/filepath"
)

func SettlementRPC() {
	repoURL := "https://github.com/airchains-network/settlement_layer_calls_api"
	destinationFolder := "dsettlement-rpc"
	destinationPath := filepath.Join(".", destinationFolder)
	if err := ChainRepoCloner(repoURL, destinationPath, "SVM"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
