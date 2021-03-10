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
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create new blog",
	Long:  `A longer description`,
	RunE: func(cmd *cobra.Command, args []string) error {
		author, err := cmd.Flags().GetString("authorID")
		if err != nil {
			return err
		}
		title, err := cmd.Flags().GetString("title")
		if err != nil {
			return err
		}
		content, err := cmd.Flags().GetString("content")
		if err != nil {
			return err
		}
		blog := &api.Blog{
			AuthorId: author,
			Title:    title,
			Content:  content,
		}

		res, err := client.CreateBlog(
			requestCtx,
			&api.CreateBlogRequest{Blog: blog},
		)
		if err != nil {
			return err
		}
		fmt.Printf("Blog created: %s\n", res.Blog)
		return nil
	},
}

func init() {
	createCmd.Flags().StringP("authorID", "a", "", "Add author ID")
	createCmd.Flags().StringP("title", "t", "", "A title for the blog")
	createCmd.Flags().StringP("content", "c", "", "The content for the blog")
	createCmd.MarkFlagRequired("author")
	createCmd.MarkFlagRequired("title")
	createCmd.MarkFlagRequired("content")
	rootCmd.AddCommand(createCmd)
}
