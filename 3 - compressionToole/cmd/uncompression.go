/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// uncompressionCmd represents the uncompression command
var uncompressionCmd = &cobra.Command{
	Use:   "ucmpr",
	Short: "ucmpr - Uncompression",
	Long:  `Command for uncompression some file`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(uncompressionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uncompressionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uncompressionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
