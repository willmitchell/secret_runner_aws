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
	"strings"
	"os/exec"
	"os"
)

type NameMap struct {
	secretName string
	envName    string
	value      string
}

func dump() {
	for i := 0; i < len(nameMap); i++ {
		m := nameMap[i]
		if verbose {
			fmt.Println(fmt.Sprintf("secretName: %v, envName: %v, value: %v", m.secretName, m.envName, m.value))
		}
	}
}

func makeEnvVarName(s string) string {
	s = strings.Replace(s, "-", "_", -1)
	s = strings.Replace(s, "/", "_", -1)
	s = strings.ToUpper(s)
	return s
}

var nameMap = []*NameMap{}
var command = ""

func addParameter(p *ssm.Parameter) {
	var nm = &NameMap{
		secretName: aws.StringValue(p.Name),
		envName:    makeEnvVarName(removeCurrentParamPath(aws.StringValue(p.Name))),
		value:      aws.StringValue(p.Value),
	}
	nameMap = append(nameMap, nm)
}

// execCmd represents the run command
var execCmd = &cobra.Command{
	Use:   "run",
	Short: "Run your program with secrets exposed as env vars.",
	Long: `The run command runs your command with environment variables where the values are set using values stored
in AWS SSM Parameter Store.  The values are decrypted using the default key within AWS KMS.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")

		sess := buildSession()

		// Create S3 service client

		cpp := getCurrentParamPath()
		fmt.Println(fmt.Sprintf("Current path: %v", cpp))

		input := ssm.GetParametersByPathInput{
			Path:           aws.String(cpp),
			Recursive:      aws.Bool(true),
			WithDecryption: aws.Bool(true),
		}

		svc := ssm.New(sess);
		o, err := svc.GetParametersByPath(&input)
		//print(o.GoString())
		if (err != nil) {
			fmt.Println(err)
			panic("Unable to get parameters by path")
		}

		params := o.Parameters
		for i := 0; i < len(params); i++ {
			p := params[i]
			//println(aws.StringValue(p.Name))
			//println(aws.StringValue(p.Value))
			//println(aws.StringValue(p.Type))
			//println(aws.Int64Value(p.Version))
			addParameter(p)
		}

		dump()
		fmt.Println("Command: ", command)

		osc := exec.Command("/bin/bash", "-c", command)
		ne := os.Environ()
		for _, cred := range nameMap {
			ne = append(ne, fmt.Sprintf("%v=%v", cred.envName, cred.value))
		}
		osc.Env = ne
		fmt.Println("Running command via bash: ", command)

		stdoutStderr, err := osc.CombinedOutput()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Println(string(stdoutStderr[:]))
	},
}
var verbose = false

func init() {
	rootCmd.AddCommand(execCmd)
	execCmd.Flags().BoolVarP(&verbose, "verbose", "", false, "Show runtime environment")
	execCmd.Flags().StringVarP(&command, "command", "c", "", "The command to run")
}
