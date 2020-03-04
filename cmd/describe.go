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

var describeActionCmd = &cobra.Command{
	Use:   "action",
	Short: "Returns the details of an action",
	Long: `Returns the details of an action
			hovercli describe action a1b2c3d4`,
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
		resp, err := APIRequest("GET", "actions/"+actionID, nil)
		if err != nil {
			log.Fatalln(err)
		}

		json.NewDecoder(resp.Body).Decode(&actionResponse)
		fmt.Print(fmt.Sprintf("ID: \t%s\nName: \t%s\nTransport Type: \t%s\nOperators: \t%v\n",
			actionResponse.Data.ID,
			actionResponse.Data.Attributes["name"],
			actionResponse.Data.Attributes["transport_type"],
			actionResponse.Data.Attributes["operators"]))
	},
}

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe resources available from the Hover API",
	Long: `Describe resources available from the Hover API
			Available subcommand is action:
				hovercli describe action a1b2c3d4`,
}

func init() {
	rootCmd.AddCommand(describeCmd)
	describeCmd.AddCommand(describeActionCmd)
}
