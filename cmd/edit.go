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
)

var editActionCmd = &cobra.Command{
	Use:   "action",
	Short: "Edit an action on Hover",
	Long: `Edit an action on Hover
	
	hover edit action a1b2c3d4 --name AirtimeBalance \
	 --root-code *144# \
	 --transport-type ussd \
	 --world-operator 2172 \
	  --world-operator 2173`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Authenticate()
		if err != nil {
			log.Fatalln(err)
		}

		if len(args) != 1 {
			log.Fatalln("Missing argument: action id")
		}

		actionID := args[0]

		var actionResponse ActionResponse
		action := ActionDetails{ 
			Name: name,
			RootCode: rootCode,
			TransportType: transportType,
			WorldOperators: worldOperators,
		}

		actionRequest := ActionRequest{CustomAction: action}
		requestBody, err := json.Marshal(actionRequest)
		fmt.Println(string(requestBody))
		resp, err := APIRequest("PATCH", "actions/"+actionID, requestBody)
		if err != nil {
			log.Fatalln(err)
		}

		json.NewDecoder(resp.Body).Decode(&actionResponse)
		fmt.Println(resp.StatusCode)
		fmt.Println("Action updated!")
		fmt.Print(fmt.Sprintf("ID: \t%s\nName: \t%s\nTransport Type: \t%s\nOperators: \t%v\n",
			actionResponse.Data.ID,
			actionResponse.Data.Attributes["name"],
			actionResponse.Data.Attributes["transport_type"],
			actionResponse.Data.Attributes["operators"]))
	},
}

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit resources on Hover",
	Long: `Edit resources on Hover

Available edit subcommands: action
hover edit action a1b2c3d4 --name AirtimeBalance`,
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.AddCommand(editActionCmd)
	editActionCmd.Flags().StringVar(&name, "name", "", "define the name of the action")
	editActionCmd.Flags().StringVar(&rootCode, "root-code", "", "define the action's root code")
	editActionCmd.Flags().StringVar(&transportType, "transport", "", "specify the action's transport type. Defaults to ussd")
	editActionCmd.Flags().StringArrayVar(&worldOperators, "world-operator", []string{}, "specify the mobile networks for the action. Multiple mobile networks can be added")
}
