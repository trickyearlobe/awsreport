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
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sso"
	"github.com/spf13/cobra"
)

// rootSsoWhatCmd represents the sso dump command
var rootSsoWhatCmd = &cobra.Command{
	Use:   "what",
	Short: "what I currently have access to in AWS through SSO",
	Run: func(cmd *cobra.Command, args []string) {

		// SSO operations need a bearer token instead of IAM credentials
		token, err := awsConfig.BearerAuthTokenProvider.RetrieveBearerToken(context.TODO())
		cobra.CheckErr(err)

		// Get ourselves an SSO client to access Identity Center stuff
		ssoClient := sso.NewFromConfig(awsConfig)

		// Iterate the accounts
		accounts, err := ssoClient.ListAccounts(context.TODO(), &sso.ListAccountsInput{AccessToken: &token.Value})
		cobra.CheckErr(err)
		for _, account := range accounts.AccountList {
			fmt.Printf("In account %s - %s\n", *account.AccountId, *account.AccountName)

			// Iterate the roles in the account
			accountRoles, err := ssoClient.ListAccountRoles(
				context.TODO(),
				&sso.ListAccountRolesInput{
					AccessToken: &token.Value,
					AccountId:   account.AccountId,
				},
			)
			cobra.CheckErr(err)
			for _, accountRole := range accountRoles.RoleList {
				fmt.Printf("    %s\n", *accountRole.RoleName)
			}
		}
	},
}

func init() {
	rootSsoCmd.AddCommand(rootSsoWhatCmd)
}
