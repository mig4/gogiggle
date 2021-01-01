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
	"os"
	"regexp"

	"github.com/spf13/cobra"

	"github.com/mig4/gogiggle/cli-experiments/cobra-k8s-selectors/pkg/cks"
	"github.com/mig4/gogiggle/cli-experiments/cobra-k8s-selectors/pkg/swapi"
)

// listCmd represents the list command options
type listCmdOptions struct {
	labelSelector string
	fieldSelector string
}

func init() {
	rootCmd.AddCommand(NewCmdList())
}

func NewCmdList() *cobra.Command {
	opts := &listCmdOptions{}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "list characters in the repo",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("list called")
			runList(opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.labelSelector, "selector", "l", opts.labelSelector, "Selector (label query) to filter characters, supports '=', '==', '!=', 'in', 'notin', 'X', '!X'.")
	flags.StringVar(&opts.fieldSelector, "field-selector", opts.fieldSelector, "Selector (field query) to filter characters, supports '=', '==', '!='.")

	return cmd
}

func runList(opts *listCmdOptions) {
	repo := swapi.NewInMemoryRepository()
	c := cks.NewCks(repo)
	characters, err := c.List(&cks.ListOptions{
		LabelSelector: opts.labelSelector,
		FieldSelector: opts.fieldSelector,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't fetch characters: %s\n", err)
		os.Exit(1)
	}

	outFormat := "%-20s %-6s %-5t %-5t\n"
	colFormat := regexp.MustCompile(`(%-?\d*)[t]`).ReplaceAllString(outFormat, "${1}s")
	fmt.Printf(colFormat, "NAME", "GENDER", "FSENS", "GHOST")
	for _, char := range characters {
		fmt.Printf(outFormat, char.Name, char.Gender, char.ForceSensitive, char.Ghost)
	}
}
