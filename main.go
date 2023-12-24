package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	commands "airchains/commands"
)

var rootCmd = &cobra.Command{
	Use:   "airchains",
	Short: "Command line tool for interacting with the Air-ZK application.",
}

var startCMD = &cobra.Command{
	Use:   "start",
	Short: "Starts the Airchains ZK-Rollup SDK by asking the user a series of questions.",
	Run:   start,
}

func start(cmd *cobra.Command, args []string) {
	commands.StartType()
}

func init() {
	rootCmd.AddCommand(startCMD)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
