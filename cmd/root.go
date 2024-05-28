/*
Copyright © 2024 Richard Nixon <richard.nixon@btinternet.com>

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
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "awsreport",
	Short: "A tool for reporting on AWS things",
	Long: `
AWS REPORT version {{}}
A tool for reporting on AWS things like users, groups, buckets etc.
The intention is to make this quite comprehensive but we will start
with identity stuff.
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var cfgFile string
var awsConfig aws.Config

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.awsreport.yaml)")
	rootCmd.PersistentFlags().String("profile", "", "AWS profile name")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".awsreport")
	}

	// Read ENV vars and config file if either exist
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	// Load AWS config as a global for all functions in "cmd" package
	err := viper.BindPFlag("profile", rootCmd.Flags().Lookup("profile"))
	cobra.CheckErr(err)
	awsDefaultConfig, loadConfigErr := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile(viper.GetString("profile")),
	)
	cobra.CheckErr(loadConfigErr)
	awsConfig = awsDefaultConfig
}
