package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"

	"github-top-n/internal/format"
	"github-top-n/internal/github"
)

var cmdTopForks = &cobra.Command{
	Use:   "forks",
	Short: "Display the top repos in an org by total forks",
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)

		ctx := context.TODO()
		github, err := github.New(githubToken, githubOrgName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		s.Suffix = " ! Loading pull request data from GitHub...this could take a while"
		s.Start()
		sortedRepos, err := github.ReposByForks(ctx)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		s.Stop()

		format.FormatForksTable(sortedRepos, totalResults, githubOrgName)
	},
}

func init() {
	rootCmd.AddCommand(cmdTopForks)
}
