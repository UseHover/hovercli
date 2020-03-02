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
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// URL points to the Hover API url
const URL = "http://localhost:3000/api/"

var cfgFile string

// Action struct represents an Action object
type Action struct {
	ID         string                 `json:"id"`
	Attributes map[string]interface{} `json:"attributes"`
}

// ActionListResponse struct represents a response containing
// a list of actions
type ActionListResponse struct {
	Data []Action
}

// ActionResponse struct represents an action response
type ActionResponse struct {
	Data Action
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hovercli",
	Short: "Welcome to the Hover Command Line Interface.",
	Long:  `Welcome to the Hover Command Line Interface.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hovercli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {

		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".hovercli" (without extension).
		viper.SetConfigName(".hovercli")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(home)

	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}

// Authenticate checks if a valid token exists. If the token is expired
// then a new one is requested
func Authenticate() error {
	authToken := viper.GetString("auth_token")
	authTokenExpiry := viper.GetTime("auth_token_expiry")

	if authToken != "" && time.Now().Before(authTokenExpiry) {
		return nil
	} else {
		var result map[string]string
		email := viper.GetString("email")
		password := viper.GetString("password")

		requestBody, err := json.Marshal(map[string]string{
			"email":    email,
			"password": password,
		})

		if err != nil {
			return err
		}

		resp, err := http.Post(URL+"authenticate", "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			return err
		}

		json.NewDecoder(resp.Body).Decode(&result)
		viper.Set("auth_token", result["auth_token"])
		viper.Set("auth_token_expiry", time.Now().Local().Add(time.Hour*2))
		err = viper.WriteConfig()
		return err
	}
}

// GetRequest makes a GET request with an Authorization header
// to the Hover API
func GetRequest(endpoint string) (*http.Response, error) {
	authToken := viper.GetString("auth_token")
	var client http.Client
	req, err := http.NewRequest("GET", URL+endpoint, nil)
	req.Header.Add("Authorization", authToken)
	if err != nil {
		return &http.Response{}, err
	}

	return client.Do(req)
}
