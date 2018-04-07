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
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list secrets for this prefix/module/stage",
	Long:`The list command shows you all of the names of the parameters in AWS SSM Parameter Store.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		sess := buildSession()

		// Create S3 service client

		cpp := getCurrentParamPath()
		fmt.Println(fmt.Sprintf("Current path: %v", cpp))
		values := []*string{aws.String(cpp)}

		pf := ssm.ParameterStringFilter{
			Key:    aws.String("Name"),
			Values: values,
			Option: aws.String("BeginsWith"),
		}

		filters := []*ssm.ParameterStringFilter{&pf}

		svc := ssm.New(sess);
		o, err := svc.DescribeParameters(&ssm.DescribeParametersInput{
			ParameterFilters: filters,
		})
		print(o.GoString())
		if (err != nil) {
			fmt.Println(err)
			panic("Unable to describe parameters")
		}

		params := o.Parameters
		for i := 0; i < len(params); i++ {
			println(aws.StringValue(params[i].Name))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
