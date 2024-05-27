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
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/spf13/cobra"
)

// usersCmd represents the users command
var rootSsoUsersCmd = &cobra.Command{
	Use:   "users",
	Short: "list of users in AWS SSO",
	Run: func(cmd *cobra.Command, args []string) {

		for _, user := range identityStoreUsers() {
			fmt.Printf("%v\n", *user.UserName)
		}
	},
}

func identityStoreUsers() []types.User {
	var userAccumulator []types.User
	for _, identityStore := range identityStores().Instances {

		// Get an identity store client for each identity store instance (normally just one)
		idStoreClient := identitystore.NewFromConfig(awsConfig)

		// Iterate all the user pages and add them to the userAccumulator
		var nextToken *string
		for morePages := true; morePages; morePages = nextToken != nil {
			usersPage, err := idStoreClient.ListUsers(context.TODO(), &identitystore.ListUsersInput{
				IdentityStoreId: identityStore.IdentityStoreId,
				NextToken:       nextToken,
			})
			cobra.CheckErr(err)
			userAccumulator = append(userAccumulator, usersPage.Users...)
			nextToken = usersPage.NextToken
		}
	}
	return userAccumulator
}

func init() {
	rootSsoCmd.AddCommand(rootSsoUsersCmd)
}
