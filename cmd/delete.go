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
	"fmt"
	"github.com/PaulKorepanow/grpcBlog/pkg/api"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long:  `A longer description`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ID, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}
		req, err := client.DeleteBlog(
			requestCtx,
			&api.DeleteBlogRequest{Id: ID},
		)
		if !req.Success {
			return fmt.Errorf("can't delete blog with id:%s", ID)
		}
		fmt.Printf("Blog with id:%v was delete", ID)
		return nil
	},
}

func init() {
	deleteCmd.Flags().StringP("id", "i", "", "The id of the blog")
	deleteCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(deleteCmd)
}
