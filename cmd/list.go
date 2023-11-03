/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	cdn "github.com/ryuichi1208/cloudfront-log-parser/lib"
	"github.com/spf13/cobra"
)

var bucket string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("list called: ", bucket)
		cdn.Get(bucket)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// bucket名をオプションで受け取る
	listCmd.Flags().StringVarP(&bucket, "bucket", "b", "default", "bucket name")
	listCmd.MarkFlagRequired("bucket")
}
