package cmd

import (
        "fmt"

        "github.com/spf13/cobra"
)

// var Version string
var commit string

func init() {
        rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
        Use:   "version",
        Short: "Print the version number of this tool",
        Long:  `All software has versions. This is used to show the version of this tool`,
        Run: func(_ *cobra.Command, _ []string) {
                fmt.Printf("commit: %s\n", commit)
        },
}
