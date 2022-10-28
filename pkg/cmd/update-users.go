package cmd

import (
	"encoding/json"
	"fmt"
	tcClient "github.com/chenliu1993/easy-tool/pkg/client"
	"github.com/chenliu1993/easy-tool/pkg/types"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// var Version string
var (
	updateFile string
	id         string
	key        string
)

func init() {
	rootCmd.AddCommand(translateUsersCmd)
	updateUsersCmd.PersistentFlags().StringVarP(&updateFile, "file", "f", "", "user account info files in json format")
	updateUsersCmd.PersistentFlags().StringVarP(&id, "id", "u", "", "cloud platform admin secret ID")
	updateUsersCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "cloud platform admin secret key")
}

var updateUsersCmd = &cobra.Command{
	Use:   "update-users",
	Short: "update-users json content in different format",
	Long:  `update-users update the userinfo directly`,
	RunE: func(_ *cobra.Command, _ []string) error {
		// read in the new version of users
		file, err := ioutil.ReadFile(updateFile)
		if err != nil {
			return err
		}
		accounts := types.Accounts{
			Data: []types.Account{},
		}
		if err := json.Unmarshal(file, &accounts); err != nil {
			return err
		}
		// generate the client
		if id == "" || key == "" {
			return fmt.Errorf("please input valid id or key")
		}
		client := tcClient.GenerateCAMClient(id, key)

		// let us do real things
		if err := tcClient.UpdateUsers(client, accounts); err != nil {
			return err
		}
		return nil
	},
}
