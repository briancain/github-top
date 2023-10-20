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

// NOTE(briancain): Ideally, these sub-commands would be setup in their own
// function rather than here in main. This was left as is for the sake of time
// given that this is an interview. But long term if I were supporting this
// project I would abstract out the commands into their own functions to be
// generated there.
func main() {
	var rootCmd = &cobra.Command{Use: "github-top"}

	var (
		// GitHub configuration
		githubOrgName string
		githubToken   string

		totalResults int
	)
	rootCmd.PersistentFlags().StringVarP(&githubToken, "token", "", "", "The GitHub token to use")
	rootCmd.PersistentFlags().StringVarP(&githubOrgName, "org-name", "o", "", "The GitHub organization to query")
	rootCmd.PersistentFlags().IntVarP(&totalResults, "total", "n", 10, "The total results to display")

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)

	ctx := context.TODO()
	var cmdTopStars = &cobra.Command{
		Use:   "stars",
		Short: "Display the top repos in an org by total stars",
		Run: func(cmd *cobra.Command, args []string) {
			github, err := github.New(githubToken, githubOrgName)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			s.Suffix = " ! Loading pull request data from GitHub...this could take a while"
			s.Start()
			sortedRepos, err := github.ReposByStars(ctx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			s.Stop()

			format.FormatStarsTable(sortedRepos, totalResults)
		},
	}
	rootCmd.AddCommand(cmdTopStars)

	var cmdTopForks = &cobra.Command{
		Use:   "forks",
		Short: "Display the top repos in an org by total forks",
		Run: func(cmd *cobra.Command, args []string) {
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

			format.FormatForksTable(sortedRepos, totalResults)
		},
	}
	rootCmd.AddCommand(cmdTopForks)

	// NOTE: this can be a slow function to run, especailly for large orgs with
	// lots of repos and lots of PRs
	var cmdTopPRs = &cobra.Command{
		Use:   "pull-requests",
		Short: "Display the top repos in an org by total pull requests",
		Run: func(cmd *cobra.Command, args []string) {
			github, err := github.New(githubToken, githubOrgName)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			s.Suffix = " ! Loading pull request data from GitHub...this could take a while"
			s.Start()
			repoPRs, err := github.ReposByPRs(ctx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			s.Stop()

			format.FormatPRsTable(repoPRs, totalResults)
		},
	}
	rootCmd.AddCommand(cmdTopPRs)

	// NOTE: this can be a slow function to run, especailly for large orgs with
	// lots of repos and lots of PRs
	var cmdTopContribs = &cobra.Command{
		Use:   "contributions",
		Short: "Display the top repos in an org by total contributions (PRs and forks)",
		Run: func(cmd *cobra.Command, args []string) {
			github, err := github.New(githubToken, githubOrgName)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			s.Suffix = " ! Loading pull request data from GitHub...this could take a while"
			s.Start()
			repoContribs, err := github.ReposByContributions(ctx)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			s.Stop()

			format.FormatContributionsTable(repoContribs, totalResults)
		},
	}
	rootCmd.AddCommand(cmdTopContribs)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
