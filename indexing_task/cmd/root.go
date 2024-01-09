package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"indexing/indexer"
)

var (
	indexFilename     string
	filenamesFilename string
)

var rootCmd = &cobra.Command{
	Use:   "indexingtask",
	Short: "Indexing task",
	Long:  `A detailed description of the indexing task command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Indexing files using", indexFilename, "and", filenamesFilename)
		if err := indexer.IndexFiles(args, indexFilename, filenamesFilename); err != nil {
			fmt.Println("Error indexing files:", err)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&indexFilename, "index", "i", "index", "Index file name")
	rootCmd.PersistentFlags().StringVarP(&filenamesFilename, "name", "n", "filenames", "Resulting file name")
}

func Execute() error {
	return rootCmd.Execute()
}
