package github

import (
	"testing"

	gh "github.com/google/go-github/v56/github"
	"github.com/stretchr/testify/require"
)

func generateTestRepositories() []*gh.Repository {
	// NOTE(briancain: The GitHub Repository struct expects its fields to be
	// var pointers. That means we can't initialize values diretly in the struct
	// and have to do this dance here and add the & when initializing a Repository
	// struct.
	name1 := "repo1"
	name2 := "repo2"
	name3 := "repo3"
	name4 := "repo4"
	name5 := "repo5"

	count1 := 10
	count2 := 1
	count3 := 50
	count4 := 0
	count5 := 100

	forkCount1 := 5
	forkCount2 := 500
	forkCount3 := 6000
	forkCount4 := 1
	forkCount5 := 23

	repos := []*gh.Repository{
		{Name: &name1, StargazersCount: &count1, ForksCount: &forkCount1},
		{Name: &name2, StargazersCount: &count2, ForksCount: &forkCount2},
		{Name: &name3, StargazersCount: &count3, ForksCount: &forkCount3},
		{Name: &name4, StargazersCount: &count4, ForksCount: &forkCount4},
		{Name: &name5, StargazersCount: &count5, ForksCount: &forkCount5},
	}

	return repos
}

func generatePRCount(repos []*gh.Repository) []RepoAndPRCount {
	var result []RepoAndPRCount

	for _, r := range repos {
		// Make PR count equal to forks for easy testing
		result = append(result, RepoAndPRCount{
			Repo:    r,
			PrCount: *r.ForksCount,
		})
	}
	return result
}

func TestSortByStars(t *testing.T) {
	repos := generateTestRepositories()

	t.Run("sorts github repositories by most stars", func(t *testing.T) {
		r := require.New(t)

		sorted := SortByStars(repos)

		r.Equal("repo5", *sorted[0].Name)
		r.Equal("repo3", *sorted[1].Name)
		r.Equal("repo1", *sorted[2].Name)
		r.Equal("repo2", *sorted[3].Name)
		r.Equal("repo4", *sorted[4].Name)
	})
}

func TestSortByForks(t *testing.T) {
	repos := generateTestRepositories()

	t.Run("sorts github repositories by most forks", func(t *testing.T) {
		r := require.New(t)

		sorted := SortByForks(repos)

		r.Equal("repo3", *sorted[0].Name)
		r.Equal("repo2", *sorted[1].Name)
		r.Equal("repo5", *sorted[2].Name)
		r.Equal("repo1", *sorted[3].Name)
		r.Equal("repo4", *sorted[4].Name)
	})
}

func TestSortByPRs(t *testing.T) {
	r := generateTestRepositories()
	repos := generatePRCount(r)

	t.Run("sorts github repositories by most pull requests", func(t *testing.T) {
		r := require.New(t)

		sorted := SortByPRs(repos)

		r.Equal("repo3", *sorted[0].Repo.Name)
		r.Equal("repo2", *sorted[1].Repo.Name)
		r.Equal("repo5", *sorted[2].Repo.Name)
		r.Equal("repo1", *sorted[3].Repo.Name)
		r.Equal("repo4", *sorted[4].Repo.Name)
	})
}

func TestSortByContributions(t *testing.T) {
	r := generateTestRepositories()
	repos := generatePRCount(r)

	t.Run("sorts github repositories by most pull requests and forks combined", func(t *testing.T) {
		r := require.New(t)

		sorted := SortByContributions(repos)

		r.Equal("repo3", *sorted[0].Repo.Name)
		r.Equal("repo2", *sorted[1].Repo.Name)
		r.Equal("repo5", *sorted[2].Repo.Name)
		r.Equal("repo1", *sorted[3].Repo.Name)
		r.Equal("repo4", *sorted[4].Repo.Name)
	})
}
