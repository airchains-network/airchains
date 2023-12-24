package client

import (
	"airchains/utils"
	"airchains/utils/evm"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

func EVMInitialiazation() {
	data := map[string]interface{}{
		"chainInfo":     EVMChainInfo(),
		"daInfo":        evmDAType(),
		"sequencerInfo": evmSequencerType(),
	}
	utils.CreateFolderAndJSONFile(data, "config", "config.json")

	daType := data["daInfo"].(map[string]interface{})["daSelected"].(string)
	daTypetoLower := strings.ToLower(daType)
	daTypeSend := "http://localhost:5050/" + daTypetoLower
	fmt.Println(daTypeSend)
	utils.ENVSetup("evm-sequencer-node", "8545", daTypetoLower)

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

		if trimmedInput == "dummyname" {
			return fmt.Errorf("please provide a custom account name other than dummyname")
		}

		return nil
	}

	prompt1 := promptui.Prompt{
		Label:    fmt.Sprintf("What do you want your account name to be? (e.g., %s)", "dummyname"),
		Validate: validate1,
	}
	accountName, err := prompt1.Run()
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

		if trimmedInput == "dummyname" || trimmedInput == "dummyName" {
			return fmt.Errorf("please provide a custom account name other than dummyname")
		}

		return nil
	}
	prompt2 := promptui.Prompt{
		Label:    fmt.Sprintf("What is the chain ID of your chain? (e.g., %s)", "aircosmic_5526-1107"),
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

		if trimmedInput == "dummyname" {
			return fmt.Errorf("please provide a custom account name other than dummyname")
		}

		return nil

	}
	prompt3 := promptui.Prompt{
		Label:    fmt.Sprintf("What do you want the name of your chain? (e.g., %s)", "dummyname"),
		Validate: validate3,
	}
	chainName, err := prompt3.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	chainInfoKeyValue = map[string]interface{}{

		"key":     accountName,
		"chainID": chainID,
		"moniker": chainName,
	}
	utils.DelayLoader()

	return chainInfoKeyValue
}

func evmDAType() map[string]interface{} {
	var chainInfoKeyValue map[string]interface{}

	prompt := promptui.Select{
		Label: "Select your DA type : ",
		Items: []string{"Celestia", "Avail", "Eigen Layer"},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | cyan }}",
			Active:   "\U000025B6 {{ if eq . \"Eigen Layer\" }}{{ . | red | cyan }} (Coming Soon){{ else }}{{ . | green }}{{ end }}",
			Inactive: "  {{ . | black }}",
			Selected: "\U000025B6 {{ . | cyan | cyan }}",
		},
	}
	_, daTech, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return evmDAType()
	}
	if daTech == "Eigen Layer" {
		fmt.Printf("Eigen Layer is coming soon. Please select other DAs for now\n")
		evmDAType()
	} else {
		fmt.Printf("You selected %s\n", daTech)
		utils.DaRPC()
		utils.SettlementRPC()
	}

	chainInfoKeyValue = map[string]interface{}{
		"daSelected": daTech,
	}

	utils.DelayLoader()
	return chainInfoKeyValue
}

func evmSequencerType() map[string]interface{} {
	var sequencerKeyValue map[string]interface{}
	prompt := promptui.Select{
		Label: "Select your sequencer type : ",
		Items: []string{"Air Sequencer", "Espresso Sequencer"},
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . | cyan }}",
			Active:   "\U000025B6 {{ if eq . \"Espresso Sequencer\" }}{{ . | red | cyan }} (Coming Soon){{ else }}{{ . | green }}{{ end }}",
			Inactive: "  {{ . | black }}",
			Selected: "\U000025B6 {{ . | cyan | cyan }}",
		},
	}
	_, sequencerValue, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		evmSequencerType()
	}
	if sequencerValue == "Espresso Sequencer" {
		fmt.Printf("Espresso Sequencer is coming soon. Only air sequencer is available for now\n")
		evmSequencerType()
	} else {
		fmt.Printf("You selected %s\n", sequencerValue)
		evm.EvmSequencerClone()
	}
	sequencerKeyValue = map[string]interface{}{
		"sequencerType": sequencerValue,
	}
	utils.DelayLoader()
	return sequencerKeyValue
}
