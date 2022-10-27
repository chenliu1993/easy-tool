package cmd

import (
	"fmt"
	"os"

	"github.com/chenliu1993/easy-tool/pkg/types"
	"github.com/spf13/cobra"
)

var (
	// This is the base resources excel file path, should be a json
	inputFile string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&inputResFile, "file", "f", "", "user account info files in json format")
}

var rootCmd = &cobra.Command{
	Use:   "easy-tool",
	Short: "easy-tool to accomploish many things",
	Long:  "easy-tool to accomploish many things",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Usually we do nothing here, see the subcommand please")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
