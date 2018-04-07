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

var encrypt = false
var param_name = ""
var param_value = ""

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "put secrets into SSM",
	Long: `The put command can store your secrets in AWS SSM Parameter Store.  These parameters may optionally
be encrypted by using the -e flag.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("put called.  name: %s, value: %s, encrypt: %t", param_name, param_value, encrypt))

		sess := buildSession()

		// Create S3 service client
		svc := ssm.New(sess);
		ev := "String"
		if (encrypt) {
			ev = "SecureString"
		}

		pi := ssm.PutParameterInput{
			Name:      aws.String(buildParamName(param_name)),
			Value:     aws.String(param_value),
			Type:      aws.String(ev),
			Overwrite: aws.Bool(true),
		}

		o, err := svc.PutParameter(&pi)
		fmt.Println(o.GoString())
		if (err != nil) {
			panic("Failed to put parameter.")
		}
	},
}

func init() {
	rootCmd.AddCommand(putCmd)

	putCmd.Flags().BoolVarP(&encrypt, "encrypt", "e", true, "Encrypt?  true/false.")

	modifyCommandWithCommonParams(putCmd)

	const PARAM_VALUE = "param_value"
	putCmd.Flags().StringVarP(&param_value, PARAM_VALUE, "v", "", "The value of the parameter")
	putCmd.MarkFlagRequired(PARAM_VALUE)


}
