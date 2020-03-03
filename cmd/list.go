/*
Copyright © 2020 Hover Developer Services <support@usehover.com>

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
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listActionsCmd = &cobra.Command{
	Use:   "actions",
	Short: "List all the actions you've configured on Hover",
	Long: `List all the actions you've configured on Hover
				hovercli list actions`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Authenticate()
		if err != nil {
			log.Fatalln(err)
		}

		var result ActionListResponse
		resp, err := GetRequest("actions?organization_id=" + viper.GetString("organization_id"))
		if err != nil {
			log.Fatalln(err)
		}

		json.NewDecoder(resp.Body).Decode(&result)

		fmt.Print(fmt.Sprintf("%-10s  %-15s  %-15s  %s\n",
			"Id", "Name", "Transport Type", "Operators"))
		for _, action := range result.Data {
			fmt.Print(fmt.Sprintf("%-10s  %-15s  %-15s  %s\n",
				action.ID,
				action.Attributes["name"],
				action.Attributes["transport_type"],
				fmt.Sprintf("%v", action.Attributes["operators"])))
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list <resource>",
	Short: "List resources (actions, transactions, parsers)",
	Long: `List resources (actions, transactions, parsers) from
			the Hover API`,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listActionsCmd)
}
