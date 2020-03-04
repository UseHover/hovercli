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

var (
	name           string
	rootCode       string
	transportType  string
	worldOperators []string
)

// ActionRequest is the request object when
// creating actions
type ActionRequest struct {
	CustomAction ActionDetails `json:"custom_action"`
}

// ActionDetails contains the Action fields
// used in ActionRequest
type ActionDetails struct {
	Name           string   `json:"name"`
	RootCode       string   `json:"root_code"`
	TransportType  string   `json:"transport_type"`
	WorldOperators []string `json:"world_operator_ids"`
}

var createActionCmd = &cobra.Command{
	Use:   "action",
	Short: "Create an action on Hover",
	Long: `Create an action on Hover
	
	hover create action --name AirtimeBalance \
	 --root-code *144# \
	 --transport-type ussd \
	 --world-operator 2172 \
	  --world-operator 2173`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Authenticate()
		if err != nil {
			log.Fatalln(err)
		}

		var actionResponse ActionResponse
		actionDetails := ActionDetails{Name: name, RootCode: rootCode,
			TransportType: transportType, WorldOperators: worldOperators}
		actionRequest := ActionRequest{CustomAction: actionDetails}
		requestBody, err := json.Marshal(actionRequest)
		resp, err := APIRequest("POST", "actions", requestBody)
		if err != nil {
			log.Fatalln(err)
		}

		json.NewDecoder(resp.Body).Decode(&actionResponse)
		fmt.Println(resp.StatusCode)
		fmt.Println("Action created!")
		fmt.Print(fmt.Sprintf("ID: \t%s\nName: \t%s\nTransport Type: \t%s\nOperators: \t%v\n",
			actionResponse.Data.ID,
			actionResponse.Data.Attributes["name"],
			actionResponse.Data.Attributes["transport_type"],
			actionResponse.Data.Attributes["operators"]))
	},
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create resources on Hover",
	Long: `Create resources on Hover

Available create subcommands: action
hover create action --name AirtimeBalance --root-code *144# --transport-type ussd --world-operator 2172`,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createActionCmd)
	createActionCmd.Flags().StringVar(&name, "name", "", "define the name of the action")
	createActionCmd.Flags().StringVar(&rootCode, "root-code", "", "define the action's root code")
	createActionCmd.Flags().StringVar(&transportType, "transport", "ussd", "specify the action's transport type. Defaults to ussd")
	createActionCmd.Flags().StringArrayVar(&worldOperators, "world-operator", []string{}, "specify the mobile networks for the action. Multiple mobile networks can be added")

	createActionCmd.MarkFlagRequired("name")
	createActionCmd.MarkFlagRequired("root-code")
	createActionCmd.MarkFlagRequired("world-operator")
}
