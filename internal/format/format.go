package format

import (
	"os"

	gh "github.com/google/go-github/v56/github"
	"github.com/jedib0t/go-pretty/v6/table"

	"github-top-n/internal/github"
)

func FormatTable(header table.Row, rows []table.Row) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(header)
	t.AppendRows(rows)
	t.SetStyle(table.StyleColoredRedWhiteOnBlack)
	t.Render()
}

func FormatStarsTable(sortedRepos []*gh.Repository, totalResults int) {
	if totalResults > len(sortedRepos) {
		totalResults = len(sortedRepos)
	}

	header := table.Row{"#", "Repo Name", "Total Stars"}
	rows := []table.Row{}
	for i := 0; i < totalResults; i++ {
		rows = append(rows, table.Row{i + 1, sortedRepos[i].GetName(), sortedRepos[i].GetStargazersCount()})
	}

	FormatTable(header, rows)
}

func FormatForksTable(sortedRepos []*gh.Repository, totalResults int) {
	if totalResults > len(sortedRepos) {
		totalResults = len(sortedRepos)
	}

	header := table.Row{"#", "Repo Name", "Total Forks"}
	rows := []table.Row{}
	for i := 0; i < totalResults; i++ {
		rows = append(rows, table.Row{i + 1, sortedRepos[i].GetName(), sortedRepos[i].GetForksCount()})
	}

	FormatTable(header, rows)
}

func FormatPRsTable(sortedRepos []github.RepoAndPRCount, totalResults int) {
	if totalResults > len(sortedRepos) {
		totalResults = len(sortedRepos)
	}

	header := table.Row{"#", "Repo Name", "Total PRs"}
	rows := []table.Row{}
	for i := 0; i < totalResults; i++ {
		rows = append(rows, table.Row{i + 1, sortedRepos[i].Repo.GetName(), sortedRepos[i].PrCount})
	}

	FormatTable(header, rows)
}

func FormatContributionsTable(sortedRepos []github.RepoAndPRCount, totalResults int) {
	if totalResults > len(sortedRepos) {
		totalResults = len(sortedRepos)
	}

	header := table.Row{"#", "Repo Name", "Total Contributions"}
	rows := []table.Row{}
	for i := 0; i < totalResults; i++ {
		rows = append(rows, table.Row{i + 1, sortedRepos[i].Repo.GetName(), sortedRepos[i].PrCount + sortedRepos[i].Repo.GetForksCount()})
	}

	FormatTable(header, rows)
}
