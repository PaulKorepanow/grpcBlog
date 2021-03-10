/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"bloggrpc/pkg/api"
	"fmt"
	"github.com/spf13/cobra"
	"io"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `A longer description `,
	RunE: func(cmd *cobra.Command, args []string) error {
		req, err := client.ListBlogs(
			requestCtx,
			&api.ListBlogRequest{},
		)
		if err != nil {
			return err
		}
		counter := 0
		for {
			res, err := req.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			fmt.Printf("%d: Blog %v\n", counter, res.GetBlog())
			counter++
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
