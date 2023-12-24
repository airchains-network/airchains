package commands

import (
	"airchains/client"
	"airchains/utils"
	"airchains/utils/cosmwasm"
	"airchains/utils/evm"
	"airchains/utils/svm"
	"fmt"

	"github.com/manifoldco/promptui"
)

func InitFunction() {
	prompt := promptui.Select{
		Label: "Select your chain type : ",
		Items: []string{"EVM", "COSMWASM", "SVM"},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | cyan }}",
			Active:   "\U000025B6 {{ . | green  }}",
			Inactive: "  {{ . | black }}",
			Selected: "\U000025B6 {{ . | cyan | cyan }}",
		},
	}

	_, chainTypeSelected, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if chainTypeSelected == "EVM" {
		evm.EVMChainClone()
		utils.DelayLoader()
		client.EVMInitialiazation()
		utils.ChainSetup("rollup-evm")
		evm.RunEVMChain()
	}

	if chainTypeSelected == "SVM" {
		svm.SolanaChainClone()
		utils.DelayLoader()
		client.SVMInitialiazation()
		svm.RunSVMChain()
	}

	if chainTypeSelected == "COSMWASM" {
		cosmwasm.CosmosChainClone()
		utils.DelayLoader()
		client.CosmWasmInitialiazation()
		utils.ChainSetup("rollup-cosmwasm")
		cosmwasm.RunCosmosChain()
	}
}
