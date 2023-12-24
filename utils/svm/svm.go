package svm

import (
	"airchains/utils"
	"fmt"
)

func RunSVMChain() {
	chainFolderName := "rollup-svm"
	chainInit := "start-chain.sh"

	seqFolderName := "svm-sequencer-node"
	seqInit := "scripts/test.sh"

	if err := utils.RunChainSeq(chainFolderName, chainInit); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if err := utils.RunChainSeq(seqFolderName, seqInit); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
