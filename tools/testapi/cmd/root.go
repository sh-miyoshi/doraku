package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
}

var rootCmd = &cobra.Command{
	Use:   "testapi",
	Short: "testapi is a command line tool for testing backend API",
	Long:  "tsearch is a command line tool for testing backend API",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute method run root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}
