package multinode

import (
	"airchains/utils"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

func MultiNodeEVMInitialiazation() {
	data := map[string]interface{}{
		"chainInfo": EVMChainInfo(),
	}
	utils.CreateFolderAndJSONFile(data, "config", "multinode-config.json")
}

func EVMChainInfo() map[string]interface{} {
	var chainInfoKeyValue map[string]interface{}

	// First question
	validate1 := func(input string) error {
		trimmedInput := strings.TrimSpace(input)

		if len(trimmedInput) == 0 {
			return fmt.Errorf("input cannot be empty")
		}

		if len(trimmedInput) < 3 {
			return fmt.Errorf("input must be at least 3 characters")
		}

		return nil
	}

	prompt1 := promptui.Prompt{
		Label:    fmt.Sprintf("What is the IP Address of the chain? (e.g., %s)", "192.168.1.1"),
		Validate: validate1,
	}
	ipAddress, err := prompt1.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	// Second question
	validate2 := func(input string) error {
		trimmedInput := strings.TrimSpace(input)

		if len(trimmedInput) == 0 {
			return fmt.Errorf("input cannot be empty")
		}

		if len(trimmedInput) < 3 {
			return fmt.Errorf("input must be at least 3 characters")
		}
		return nil
	}
	prompt2 := promptui.Prompt{
		Label:    fmt.Sprintf("What is the chain ID of your chain? (e.g., %s)", "aircosmic_5501-1107"),
		Validate: validate2,
	}
	chainID, err := prompt2.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	// Third question
	validate3 := func(input string) error {
		trimmedInput := strings.TrimSpace(input)

		if len(trimmedInput) == 0 {
			return fmt.Errorf("input cannot be empty")
		}

		if len(trimmedInput) < 3 {
			return fmt.Errorf("input must be at least 3 characters")
		}
		return nil

	}
	prompt3 := promptui.Prompt{
		Label:    fmt.Sprintf("What is the node ID of your chain? (e.g., %s)", "d27015f43a1b5de64ecda0f99ac00a0de34f3d95"),
		Validate: validate3,
	}
	nodeID, err := prompt3.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	// Second question
	validate4 := func(input string) error {
		trimmedInput := strings.TrimSpace(input)

		if len(trimmedInput) == 0 {
			return fmt.Errorf("input cannot be empty")
		}

		if len(trimmedInput) < 3 {
			return fmt.Errorf("input must be at least 3 characters")
		}

		return nil
	}
	prompt4 := promptui.Prompt{
		Label:    fmt.Sprintf("What do you want to name your node? (e.g., %s)", "aircosmic-node-1"),
		Validate: validate4,
	}
	nodeName, err := prompt4.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	prompt5 := promptui.Prompt{
		Label:    fmt.Sprintf("More Info at https://docs.airchains.io  Paste the seed node genesis.json at evm-chain/seed-genesis.json. Type Done and Press Enter to continue"),
		Validate: validate4,
	}
	genesisFile, err := prompt5.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	chainInfoKeyValue = map[string]interface{}{
		"ipAddress": ipAddress,
		"chainID":   chainID,
		"nodeID":    nodeID,
		"nodeName":  nodeName,
		"genesis":   genesisFile,
	}
	utils.DelayLoader()

	return chainInfoKeyValue
}
