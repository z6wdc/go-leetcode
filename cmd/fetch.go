package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/z6wdc/go-leetcode/internal/usecase"
	"os"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch <slug>",
	Short: "Fetch a LeetCode problem by slug and generate code/test/README files",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		slug := args[0]
		err := usecase.GenerateQuestionBySlug(slug)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to generate question: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Successfully generated files for %s\n", slug)
	},
}
