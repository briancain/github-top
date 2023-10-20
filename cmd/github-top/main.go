package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "github-top",
	Short:   "A command line tool to interact with GitHub organizations",
	Example: "github-top stars --org-name netflix",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&githubToken, "token", "", "", "The GitHub token to use")
	rootCmd.PersistentFlags().StringVarP(&githubOrgName, "org-name", "o", "", "The GitHub organization to query")
	rootCmd.PersistentFlags().IntVarP(&totalResults, "total", "n", 10, "The total results to display")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
