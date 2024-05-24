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
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/spf13/cobra"
)

// rootWhoCmd represents the who command
var rootWhoCmd = &cobra.Command{
	Use:   "who",
	Short: "get your STS identity",
	Long:  `fetches your STS identity for the selected or default AWS profile`,
	Run: func(cmd *cobra.Command, args []string) {

		stsClient := sts.NewFromConfig(awsConfig)
		identity, getCallerIdErr := stsClient.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
		cobra.CheckErr(getCallerIdErr)

		fmt.Printf("AWS STS UserID:  %v\n", aws.ToString(identity.UserId))
		fmt.Printf("AWS STS Account: %v\n", aws.ToString(identity.Account))
		fmt.Printf("AWS STS ARN:     %v\n", aws.ToString(identity.Arn))
	},
}

func init() {
	rootCmd.AddCommand(rootWhoCmd)
}
