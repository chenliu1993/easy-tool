package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/chenliu1993/easy-tool/pkg/types"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// var Version string
var (
	format    string
	output    string
	inputFile string
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().StringVarP(&inputFile, "file", "f", "", "user account info files in json format")
	getCmd.PersistentFlags().StringVarP(&format, "format", "", "csv", "output format")
	getCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "output to un-stdout")
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get json content in different format",
	Long:  `get json content in different format`,
	RunE: func(_ *cobra.Command, _ []string) error {
		if err := translate(inputFile, output); err != nil {
			return err
		}
		return nil
	},
}

func translate(in, out string) error {
	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}
	accounts := types.Accounts{
		Data: []types.Account{},
	}
	if err := json.Unmarshal(file, &accounts); err != nil {
		return err
	}

	// default csv
	if out != "" {
		// not to stdout
		outputFile, err := os.OpenFile(out, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		if _, err := outputFile.Write([]byte("ConsoleLogin,CountryCode,CreateTime,Email,Name,NickName,PhoneNum,Remark,Uid,Uin\n")); err != nil {
			return err
		}

		for _, a := range accounts.Data {
			line := fmt.Sprintf("%d,%s,%s,%s,%s,%s,%s,%s,%d,%d\n",
				a.ConsoleLogin, a.CountryCode, a.CreateTime, a.Email, a.Name, a.NickName, a.PhoneNum, a.Remark, a.Uid, a.Uin)
			if _, err := outputFile.Write([]byte(line)); err != nil {
				return err
			}
		}

		return nil
	}
	for _, a := range accounts.Data {
		line := fmt.Sprintf("ConsoleLogin:%d,CountryCode:%s,CreateTime:%s,Email:%s,Name:%s,NickName:%s,PhoneNum:%s,Remark:%s,Uid:%d,Uin:%d\n",
			a.ConsoleLogin, a.CountryCode, a.CreateTime, a.Email, a.Name, a.NickName, a.PhoneNum, a.Remark, a.Uid, a.Uin)
		fmt.Println(line)
	}

	return nil
}
