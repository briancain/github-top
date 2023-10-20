package format

import (
	"fmt"
	"os"

	gh "github.com/google/go-github/v56/github"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github-top-n/internal/github"
)

func FormatTable(header table.Row, rows []table.Row, title string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(header)
	t.AppendRows(rows)
	t.SetStyle(table.StyleColoredRedWhiteOnBlack)
	t.SetTitle(title)
	t.Style().Title.Align = text.AlignCenter
	t.Render()
}

func FormatStarsTable(sortedRepos []*gh.Repository, totalResults int, org string) {
	if totalResults > len(sortedRepos) {
		totalResults = len(sortedRepos)
	}

	header := table.Row{"#", "Repo Name", "Total Stars"}
	rows := []table.Row{}
	for i := 0; i < totalResults; i++ {
		rows = append(rows, table.Row{i + 1, sortedRepos[i].GetName(), sortedRepos[i].GetStargazersCount()})
	}

	title := fmt.Sprintf("Top %d Stars for %q", totalResults, org)

	FormatTable(header, rows, title)
}

func FormatForksTable(sortedRepos []*gh.Repository, totalResults int, org string) {
	if totalResults > len(sortedRepos) {
		totalResults = len(sortedRepos)
	}

	header := table.Row{"#", "Repo Name", "Total Forks"}
	rows := []table.Row{}
	for i := 0; i < totalResults; i++ {
		rows = append(rows, table.Row{i + 1, sortedRepos[i].GetName(), sortedRepos[i].GetForksCount()})
	}

	title := fmt.Sprintf("Top %d Forks for %q", totalResults, org)

	FormatTable(header, rows, title)
}

func FormatPRsTable(sortedRepos []github.RepoAndPRCount, totalResults int, org string) {
	if totalResults > len(sortedRepos) {
		totalResults = len(sortedRepos)
	}

	header := table.Row{"#", "Repo Name", "Total PRs"}
	rows := []table.Row{}
	for i := 0; i < totalResults; i++ {
		rows = append(rows, table.Row{i + 1, sortedRepos[i].Repo.GetName(), sortedRepos[i].PrCount})
	}

	title := fmt.Sprintf("Top %d Pull Requests for %q", totalResults, org)

	FormatTable(header, rows, title)
}

func FormatContributionsTable(sortedRepos []github.RepoAndPRCount, totalResults int, org string) {
	if totalResults > len(sortedRepos) {
		totalResults = len(sortedRepos)
	}

	header := table.Row{"#", "Repo Name", "Total Contributions"}
	rows := []table.Row{}
	for i := 0; i < totalResults; i++ {
		rows = append(rows, table.Row{i + 1, sortedRepos[i].Repo.GetName(), sortedRepos[i].PrCount + sortedRepos[i].Repo.GetForksCount()})
	}

	title := fmt.Sprintf("Top %d Contributions (PRs + Forks) for %q", totalResults, org)

	FormatTable(header, rows, title)
}
