package commands

import (
	"airchains/client/multinode"
	"airchains/utils"
	"airchains/utils/evm"
	"fmt"

	"github.com/manifoldco/promptui"
)

func RunNonValidatorNode() {
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
		multinode.MultiNodeEVMInitialiazation()
		evm.RunMultiNodeEVMChain()
	}
}
