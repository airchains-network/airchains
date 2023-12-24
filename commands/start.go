package commands

import (
	"airchains/utils"
	"fmt"

	"github.com/manifoldco/promptui"
)

func StartType() {
	prompt := promptui.Select{
		Label: "Select your chain type : ",
		Items: []string{"Start a new chain", "Run a non validator node", "Run a validator node"},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | cyan }}",
			Active:   "\U000025B6 {{ . | green  }}",
			Inactive: "  {{ . | black }}",
			Selected: "\U000025B6 {{ . | cyan | cyan }}",
		},
	}

	_, selectedStartType, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if selectedStartType == "Start a new chain" {
		fmt.Println("Starting a new chain...")
		utils.DelayLoader()
		InitFunction()

	}

	if selectedStartType == "Run a non validator node" {
		fmt.Println("Running a non validator node...")
		fmt.Println("Please refer to the documentation at https://docs.airchains.io for more information")
		utils.DelayLoader()
		RunNonValidatorNode()
	}
}
