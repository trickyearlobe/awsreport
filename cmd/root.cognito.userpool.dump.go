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

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "dump AWS Cognito user pool group memberships",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CognitoPool,CognitoGroup,CognitoUserName,CognitoUserEmail")
		for _, pool := range cognitoUserPools() {
			for _, group := range cognitoUserPoolGroups(*pool.Id) {
				for _, member := range cognitoUserPoolGroupMembers(*pool.Id, *group.GroupName) {
					fmt.Printf("%v,%v,%v,%v\n", *pool.Name, *group.GroupName, *member.Username, cognitoUserEmail(*member.Username, *pool.Id))
				}
			}
		}
	},
}

func cognitoUserEmail(userName string, poolId string) string {
	cognitoIdentityClient := cognitoidentityprovider.NewFromConfig(awsConfig)
	userDetails, err := cognitoIdentityClient.AdminGetUser(context.TODO(), &cognitoidentityprovider.AdminGetUserInput{
		Username:   &userName,
		UserPoolId: &poolId,
	})
	cobra.CheckErr(err)
	emailAddr := ""
	for _, attribute := range userDetails.UserAttributes {
		if *attribute.Name == "email" {
			emailAddr = *attribute.Value
		}
	}
	return emailAddr
}

func cognitoUserPoolGroupMembers(userPoolID string, groupId string) []types.UserType {
	var groupMembersAccumulator []types.UserType
	cognitoIdentityClient := cognitoidentityprovider.NewFromConfig(awsConfig)
	var nextToken *string
	for morePages := true; morePages; morePages = nextToken != nil {
		members, err := cognitoIdentityClient.ListUsersInGroup(context.TODO(), &cognitoidentityprovider.ListUsersInGroupInput{
			UserPoolId: &userPoolID,
			GroupName:  &groupId,
			NextToken:  nextToken,
		})
		cobra.CheckErr(err)
		groupMembersAccumulator = append(groupMembersAccumulator, members.Users...)
		nextToken = members.NextToken
	}
	return groupMembersAccumulator
}

func cognitoUserPoolGroups(userPoolId string) []types.GroupType {
	var groupsAccumulator []types.GroupType
	cognitoIdentityClient := cognitoidentityprovider.NewFromConfig(awsConfig)
	var nextToken *string
	for morePages := true; morePages; morePages = nextToken != nil {
		groups, err := cognitoIdentityClient.ListGroups(context.TODO(), &cognitoidentityprovider.ListGroupsInput{
			UserPoolId: &userPoolId,
			NextToken:  nextToken,
		})
		cobra.CheckErr(err)
		groupsAccumulator = append(groupsAccumulator, groups.Groups...)
		nextToken = groups.NextToken
	}
	return groupsAccumulator
}

func init() {
	rootCognitoUserPoolCmd.AddCommand(dumpCmd)
}
