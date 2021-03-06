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

// getCmd represents the put command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get secrets from SSM",
	Long:`The get command can be used to show what parameters you have in AWS SSM Parameter Store.
This can be useful when trying to figure out how the name/path hierarchy works in SSM PS.`,
	Run: func(cmd *cobra.Command, args []string) {

		sess := buildSession()

		svc := ssm.New(sess);
		name := buildParamName(param_name)
		if !Raw_output{
			fmt.Println(fmt.Sprintf("Getting parameter: name: %s, computed path: %s", param_name, name))
		}
		pi := ssm.GetParameterInput{
			Name:      aws.String(name),
			WithDecryption: aws.Bool(true),
		}

		o, err := svc.GetParameter(&pi)
		if !Raw_output{
			fmt.Print(o.GoString())
		} else {
			fmt.Println(aws.StringValue(o.Parameter.Value))
		}
		if (err != nil) {
			fmt.Println("Failed to get parameter.",err)
			os.Exit(-1)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	modifyCommandWithCommonParams(getCmd)
}

var Raw_output = false

func modifyCommandWithCommonParams(cmd *cobra.Command) {
	const PARAM_NAME = "param_name"
	cmd.Flags().StringVarP(&param_name, PARAM_NAME, "n", "", "The name of the parameter")
	cmd.MarkFlagRequired(PARAM_NAME)
	const RAW_OUTPUT = "raw_output"
	cmd.Flags().BoolVarP(&Raw_output, RAW_OUTPUT, "", false, "Provide secret output in raw format")
}
