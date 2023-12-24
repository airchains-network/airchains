package evm

import (
	"airchains/utils"
	"fmt"
)

func RunEVMChain() {
	chainFolderName := "rollup-evm"
	chainInit := "start-chain.sh"

	seqFolderName := "evm-sequencer-node"
	seqInit := "scripts/test.sh"

	if err := utils.RunChainSeq(chainFolderName, chainInit); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if err := utils.RunChainSeq(seqFolderName, seqInit); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
