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
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/spf13/cobra"
)

// rootSsoCmd represents the sso command
var rootSsoCmd = &cobra.Command{
	Use:   "sso",
	Short: "commands for AWS Identity Manager SSO",
}

func identityStores() *ssoadmin.ListInstancesOutput {
	ssoAdminClient := ssoadmin.NewFromConfig(awsConfig)
	ids, err := ssoAdminClient.ListInstances(context.TODO(), &ssoadmin.ListInstancesInput{})
	cobra.CheckErr(err)
	return ids
}

func init() {
	rootCmd.AddCommand(rootSsoCmd)
}
