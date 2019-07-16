/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		jsonString, err := getRegistry()
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", jsonString)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&regServer, "server", "s", "", "server hosting registry")
	listCmd.MarkFlagRequired("server")
}

func getRegistry() (jsonString string, err error) {

	resp, err := http.Get(fmt.Sprintf("http://%s", regServer))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// the registry (the response) -> {host1:{hostname:"hostname", remote_addr:"123.123.123.123"...etc}, host2:{hostname:"hostname", remote_addr:"123.123.123.234"...etc}}
	// respBody := make(map[string]map[string]string)
	// decoder := json.NewDecoder(resp.Body)
	// decodeErr := decoder.Decode(&respBody)
	// if decodeErr != nil {
	// 	return "", decodeErr
	// }
	// jsonBytes, marshErr := json.MarshalIndent(respBody, "", "  ")
	// if marshErr != nil {
	// 	return "", marshErr
	// }

	return string(body), nil

}
