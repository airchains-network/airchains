package cosmwasm

import (
	"airchains/utils"
	"fmt"
)

func RunCosmosChain() {
	chainFolderName := "rollup-cosmwasm"
	chainInit := "start-chain.sh"

	seqFolderName := "cosmwasm-sequencer-node"
	seqInit := "scripts/test.sh"

	if err := utils.RunChainSeq(chainFolderName, chainInit); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if err := utils.RunChainSeq(seqFolderName, seqInit); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
