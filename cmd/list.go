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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		sess := buildSession()

		// Create S3 service client
		svc := ssm.New(sess);
		o, err := svc.DescribeParameters(&ssm.DescribeParametersInput{})
		if (err != nil) {
			panic("Unable to describe parameters")
		}
		//print(o.GoString())

		params := o.Parameters
		for i := 0; i < len(params); i++ {
			println(aws.StringValue(params[i].Name))
		}
		println(module)
		println(stage)
		println(buildParamName(prefix,module,stage,"hello"))

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}