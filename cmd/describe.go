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
	"fmt"
	"log"
	"encoding/json"

	"github.com/spf13/cobra"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe an action",
	Long: `Describe an action
			hovercli describe a1b2c3d4`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Authenticate()
		if err != nil {
			log.Fatalln(err)
		}

		if len(args) != 1 {
			log.Fatalln("Missing argument: action id")
		}

		actionId := args[0]
		var result map[string]interface{}
		resp, err := GetRequest("actions/" + actionId)
		if err != nil {
			log.Fatalln(err)
		}
		
		json.NewDecoder(resp.Body).Decode(&result)
		fmt.Println(result["data"])
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)
}
