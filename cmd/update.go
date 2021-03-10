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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ID, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		authorID, err := cmd.Flags().GetString("author_id")
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
		req, err := client.UpdateBlog(
			requestCtx,
			&api.UpdateBlogRequest{Blog: &api.Blog{
				Id:       ID,
				AuthorId: authorID,
				Title:    title,
				Content:  content,
			}},
		)
		if err != nil {
			return err
		}
		fmt.Printf("Blog is updated:%v", req.GetBlog())
		return nil
	},
}

func init() {
	updateCmd.Flags().StringP("id", "i", "", "ID of the blog")
	updateCmd.Flags().StringP("author_id", "a", "", "ID of the author")
	updateCmd.Flags().StringP("title", "t", "", "A title for the blog")
	updateCmd.Flags().StringP("content", "c", "", "The content for the blog")
	updateCmd.MarkFlagRequired("id")
	updateCmd.MarkFlagRequired("author_id")
	updateCmd.MarkFlagRequired("title")
	updateCmd.MarkFlagRequired("content")
	rootCmd.AddCommand(updateCmd)
}
