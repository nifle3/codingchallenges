/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// compressionCmd represents the compression command
var compressionCmd = &cobra.Command{
	Use:   "cmpr",
	Short: "To comprese some file",
	Long:  `Use for compression some file. This is a third quest from coding challenges`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println(string(colorRed) + "Please enter a file path" + string(colorReset))
			os.Exit(0)
		}

		fileName := args[0]

		_, err := os.Stat(fileName)
		if os.IsNotExist(err) {
			fmt.Println(colorRed + "Your file is not exist" + colorReset)
			os.Exit(0)
		}

		fmt.Printf("File exist\n")

		file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
		if err != nil {
			fmt.Println(colorRed, err.Error(), colorReset)
			os.Exit(0)
		}
		defer func() {
			if err := file.Close(); err != nil {
				fmt.Println(colorRed, err.Error(), colorReset)
				os.Exit(0)
			}
		}()

		scanner := bufio.NewScanner(file)
		numberOfChar := make(map[byte]int)
		for scanner.Scan() {
			line := scanner.Text()

			for idx := range line {
				numberOfChar[line[idx]] += 1
			}
		}

		for key, value := range numberOfChar {
			fmt.Printf("%s : %d\n", string(key), value)
		}
	},
}

func init() {
	rootCmd.AddCommand(compressionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compressionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compressionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
