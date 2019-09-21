package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var userConfig = UserConfig{}

func init() {
	userAddCmd.Flags().StringVarP(&userConfig.Name, "name", "n", "", "a name of new user")
	userAddCmd.MarkFlagRequired("name")
	userAddCmd.Flags().StringVarP(&userConfig.Password, "password", "p", "", "a password of new user")

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
		// TODO(Input Password)

		url := globalConfig.BackendServerAddr + "/api/v1"
		body := fmt.Sprintf(`{"name": "%s","password": "%s"}`, userConfig.Name, userConfig.Password)
		res, err := http.Post(url+"/user", "application/json", bytes.NewBufferString(body))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to request server: %v\n", err)
			return
		}

		switch res.StatusCode {
		case 202:
			fmt.Println("Success to get validate code")
		case 400:
			fmt.Fprintf(os.Stderr, "Failed with Missing Request: Name[%s], Password[%s]\n", userConfig.Name, userConfig.Password)
			return
		default:
			fmt.Fprintf(os.Stderr, "Unexpected Response form Server: %s", res.Status)
			return
		}

		// TODO(validate user)
	},
}
