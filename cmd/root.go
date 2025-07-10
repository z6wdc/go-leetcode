package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "leetcode",
	Short: "CLI tool for working with LeetCode problems",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
