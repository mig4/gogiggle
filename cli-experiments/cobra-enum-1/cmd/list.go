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

	"github.com/creachadair/goflags/enumflag"
	"github.com/spf13/cobra"
)

var (
	status = enumflag.New("", "ACTIVE", "DISABLED", "EXPIRED")

	// listCmd represents the list command
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("list called")
			fmt.Printf("status: key: %s, str: %s, idx: %d, get: %v\n", status.Key(), status.String(), status.Index(), status.Get())
			if fs, err := Parse(status.Key()); err == nil {
				fmt.Printf("filter status: %+v\n", fs)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)

	// To use `enumflag.Value` with pflag we need to wrap it in a compatibility
	// interface that provides `Type()` method.
	// The `status.Help(string)` method returns the given help string with the
	// allowed values appended to it.
	listCmd.Flags().Var(CompatValue{wrapped: status}, "status", status.Help("filter for given status"))
}
