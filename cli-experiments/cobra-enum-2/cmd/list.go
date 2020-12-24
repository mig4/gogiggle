/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag"
)

var fooStatus FooStatus

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		fsStr := cmd.Flags().Lookup("status").Value.String()
		fmt.Printf("fooStatus: %+v\n", fooStatus)
		fmt.Printf("fsStr: %+v\n", fsStr)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().Var(
		enumflag.New(&fooStatus, "status", FooStatusIds, enumflag.EnumCaseInsensitive),
		"status",
		"filter for given status",
	)
}
