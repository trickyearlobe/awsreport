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
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/spf13/cobra"
)

// rootCognitoUserPoolsListCmd represents the pools command
var rootCognitoUserPoolListCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all AWS Cognito User Pools",
	Run: func(cmd *cobra.Command, args []string) {
		for _, pool := range cognitoUserPools() {
			fmt.Printf("%v (%v)\n", *pool.Name, *pool.Id)
		}
	},
}

func cognitoUserPools() []types.UserPoolDescriptionType {
	identityClient := cognitoidentityprovider.NewFromConfig(awsConfig)
	var userPoolAccumulator []types.UserPoolDescriptionType
	maxResults := int32(60)
	var nextToken *string
	for morePages := true; morePages; morePages = nextToken != nil {
		userPools, err := identityClient.ListUserPools(context.TODO(), &cognitoidentityprovider.ListUserPoolsInput{
			MaxResults: &maxResults,
			NextToken:  nextToken,
		})
		cobra.CheckErr(err)
		userPoolAccumulator = append(userPoolAccumulator, userPools.UserPools...)
		nextToken = userPools.NextToken
	}
	return userPoolAccumulator
}

func init() {
	rootCognitoUserPoolCmd.AddCommand(rootCognitoUserPoolListCmd)
}
