// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/aws"
	"os"
)

// deleteCmd represents the put command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete secrets from SSM",
	Long:`The delete command can be used to remove secrets from AWS SSM Parameter Store.`,
	Run: func(cmd *cobra.Command, args []string) {

		sess := buildSession()

		svc := ssm.New(sess);
		name := buildParamName(param_name)
		fmt.Println(fmt.Sprintf("Getting parameter: name: %s, computed path: %s", param_name, name))
		pi := ssm.DeleteParameterInput{
			Name:      aws.String(name),
		}

		o, err := svc.DeleteParameter(&pi)
		fmt.Println(o.GoString())
		if (err != nil) {
			fmt.Println("Failed to delete parameter.",err)
			os.Exit(-1)
		}
		fmt.Println("Parameter deleted")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	modifyCommandWithCommonParams(deleteCmd)
}
