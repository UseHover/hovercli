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
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured actions",
	Long: `List configured actions`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Authenticate()
		if err != nil {
			log.Fatalln(err)
		}

		var result map[string]interface{}
		resp, err := GetRequest("actions?organization_id="+viper.GetString("organization_id"))
		if err != nil {
			log.Fatalln(err)
		}

		json.NewDecoder(resp.Body).Decode(&result)
		fmt.Println(result["data"])
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}