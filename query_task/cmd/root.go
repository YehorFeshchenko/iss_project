package cmd

import (
	"github.com/spf13/cobra"
	"query/query"
)

var (
	indexFilename     string
	filenamesFilename string
)

var rootCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the index",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		word := args[0]
		query.ExecuteQuery(word, indexFilename, filenamesFilename)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&indexFilename, "index", "i", "index", "Index file name")
	rootCmd.PersistentFlags().StringVarP(&filenamesFilename, "name", "n", "filenames", "Filenames file name")
}

func Execute() error {
	return rootCmd.Execute()
}
