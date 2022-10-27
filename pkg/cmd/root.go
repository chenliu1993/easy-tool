package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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
