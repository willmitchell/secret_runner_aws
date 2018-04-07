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
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string
var prefix = ""
var module = ""
var stage = ""

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "secret_runner_aws",
	Short: "This program helps you run your own programs in a secure manner by passing secrets as env vars.",
	Long:
	`This program helps you use AWS SSM Parameter Store to manage your parameters and secrets.  These secrets are
encrypted using master keys that are managed by AWS KMS.  The tool provides CRUD operations for params 
and secrets, and it relies on a naming convention that maps onto existing AWS SSM PS IAM role management 
facilities.  You can also use this program to run *your* program in a secure manner using the 'run' command.
The run command exposes all appropriate secrets as environment variables and then execs your program.  

Disclaimer:   Provided without warranty of any kind.  Use at your own risk.  
Bug reports:  https://gitlab.com/willmitchell/secret_runner_aws/issues
Author:       will.mitchell@app3.com, 2018.
`,

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.secret_runner_aws.yaml)")

	const PREFIX = "prefix"
	rootCmd.PersistentFlags().StringVarP(&prefix, PREFIX, "p", "", "prefix name", )
	rootCmd.MarkFlagRequired(PREFIX)
	const MODULE = "module"
	rootCmd.PersistentFlags().StringVarP(&module, MODULE, "m", "", "Module name", )
	rootCmd.MarkFlagRequired(MODULE)
	const STAGE = "stage"
	rootCmd.PersistentFlags().StringVarP(&stage, STAGE, "s", "", "Stage name", )
	rootCmd.MarkFlagRequired(STAGE)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//initConfig reads in config file and ENV variables if set.
func initConfig() {
	//if cfgFile != "" {
	//	// Use config file from the flag.
	//	//viper.SetConfigFile(cfgFile)
	//} else {
	//	//// Find home directory.
	//	//home, err := homedir.Dir()
	//	//if err != nil {
	//	//	fmt.Println(err)
	//	//	os.Exit(1)
	//	//}
	//
	//	//// Search config in home directory with name ".secret_runner_aws" (without extension).
	//	//viper.AddConfigPath(home)
	//	//viper.SetConfigName(".secret_runner_aws")
	//}
	//
	//viper.AutomaticEnv() // read in environment variables that match
	//
	//// If a config file is found, read it in.
	//if err := viper.ReadInConfig(); err == nil {
	//	fmt.Println("Using config file:", viper.ConfigFileUsed())
	//}
}
