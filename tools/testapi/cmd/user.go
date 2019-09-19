package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var userConfig = UserConfig{}

func init() {
	userCmd.Flags().StringVarP(&userConfig.Name, "name", "n", "", "a name of new user")
	userCmd.MarkFlagRequired("name")

	userCmd.AddCommand(userAddCmd)

	rootCmd.AddCommand(userCmd)
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "manage user resource",
	Long:  `manage user resource`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Println("user command requires subcommand")
	},
}

var userAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add new user",
	Long:  `add new user`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new user")
	},
}
