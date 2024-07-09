/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	colorRed   = "\033[31m"
	colorReset = "\033[0m"
)

var rootCmd = &cobra.Command{
	Use:   "compressinTool",
	Short: "The tool for compression and uncompression some files",
	Long:  `This application has been created for solve coding challenge three quest`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
