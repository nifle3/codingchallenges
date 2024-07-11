/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nifle3/compressinTool/pkg/dataStruct"
	"github.com/nifle3/compressinTool/pkg/output"

	"github.com/spf13/cobra"
)

// compressionCmd represents the compression command
var compressionCmd = &cobra.Command{
	Use:   "cmpr",
	Short: "To comprese some file",
	Long:  `Use for compression some file. This is a third quest from coding challenges`,
	Run: func(cmd *cobra.Command, args []string) {
		out := output.CreateOutput()

		if len(args) < 1 {
			out.Error("You must enter a file path")
			os.Exit(0)
		}

		fileName := args[0]

		_, err := os.Stat(fileName)
		if os.IsNotExist(err) {
			out.Error("Your file is not exist")
			os.Exit(0)
		}

		fmt.Printf("File exist\n")

		file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
		if err != nil {
			out.Error(err.Error())
			os.Exit(0)
		}
		defer func() {
			if err := file.Close(); err != nil {
				out.Error(err.Error())
				os.Exit(0)
			}
		}()

		scanner := bufio.NewScanner(file)
		scanner.Buffer(make([]byte, 1024*1024), bufio.MaxScanTokenSize)
		numberOfChar := make(map[byte]int)
		for scanner.Scan() {
			line := scanner.Text()

			for idx := range line {
				numberOfChar[line[idx]] += 1
			}
		}

		queue := dataStruct.New(len(numberOfChar))

		for key, value := range numberOfChar {
			node := dataStruct.Node{}
			node.Elem = rune(key)
			node.Freq = value

			queue.Insert(node)

			fmt.Printf("%s : %d\n", string(key), value)
		}

		var lastNode dataStruct.Node
		for queue.Length() > 0 {
			node1, err := queue.ExtractMinimum()
			if err != nil {
				break
			}

			node2, err := queue.ExtractMinimum()
			if err != nil {
				lastNode = node1
				break
			}

			node := dataStruct.Node{}
			node.Left = &node1
			node.Right = &node2
			node.Freq = node1.Freq + node2.Freq

			queue.Insert(node)
		}

		lastNode.PrintTree()
	},
}

func init() {
	rootCmd.AddCommand(compressionCmd)
}
