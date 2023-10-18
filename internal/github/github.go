package github

import (
	"context"
	"errors"
	"sort"

	gh "github.com/google/go-github/v56/github"
)

type GitHub struct {
	Client *gh.Client

	Org string
}

func New(accessToken string, org string) (*GitHub, error) {
	if org == "" {
		return nil, errors.New("No GitHub organization specified")
	}

	var c *gh.Client
	if accessToken == "" {
		c = gh.NewClient(nil)
	} else {
		c = gh.NewClient(nil).WithAuthToken(accessToken)
	}

	return &GitHub{
		Client: c,
		Org:    org,
	}, nil
}

// FetchRepos uses the Github API to query for all repositories by a given Organization.
// It will iterate through the entire set of repos and respect GitHubs pagination.
func (g *GitHub) FetchRepos(ctx context.Context) ([]*gh.Repository, error) {
	opt := &gh.RepositoryListByOrgOptions{
		Type:        "all",
		ListOptions: gh.ListOptions{PerPage: 100},
	}
	repositories := make([]*gh.Repository, 0)

	// Iterate through all results and respect GH pagination
	for {
		repos, resp, err := g.Client.Repositories.ListByOrg(context.Background(), g.Org, opt)
		if err != nil {
			return nil, err
		}
		repositories = append(repositories, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return repositories, nil
}

// FetchPRCount will take a repo string, and look up that repo based on the org
// given and return the total number of ALL pull requests for that repo.
func (g *GitHub) FetchPRCount(ctx context.Context, repo string) (int, error) {
	opt := &gh.PullRequestListOptions{
		State:       "all", // all of the PRs
		ListOptions: gh.ListOptions{PerPage: 100},
	}

	pullRequests := []*gh.PullRequest{}
	// Iterate through all results and respect GH pagination
	for {
		prs, response, err := g.Client.PullRequests.List(ctx, g.Org, repo, opt)
		if err != nil {
			return 0, err
		}
		pullRequests = append(pullRequests, prs...)
		if response.NextPage == 0 {
			break
		}
		opt.Page = response.NextPage
	}

	return len(pullRequests), nil
}

// NOTE: we create a temporary struct to hold onto the PR count because the
// GitHub golang API does not include that on its Repository struct.
type RepoAndPRCount struct {
	Repo    *gh.Repository
	PrCount int
}

// CollectRepoPRCount takes a list of github repositories and queries for all
// pull requests via the GitHub API. It returns a struct of the repo and its PR
// total count
func (g *GitHub) CollectRepoPRCount(
	ctx context.Context,
	repositories []*gh.Repository,
) ([]RepoAndPRCount, error) {
	var result []RepoAndPRCount
	for _, r := range repositories {
		total, err := g.FetchPRCount(ctx, r.GetName())
		if err != nil {
			return nil, err
		}

		result = append(result, RepoAndPRCount{
			Repo:    r,
			PrCount: total,
		})
	}

	return result, nil
}

func (g *GitHub) ReposByStars(
	ctx context.Context,
) ([]*gh.Repository, error) {
	repositories, err := g.FetchRepos(ctx)
	if err != nil {
		return nil, err
	}

	result := SortByStars(repositories)

	return result, nil
}

func SortByStars(repositories []*gh.Repository) []*gh.Repository {
	// Sort by stars
	sort.Slice(repositories, func(i, j int) bool {
		return repositories[i].GetStargazersCount() > repositories[j].GetStargazersCount()
	})

	return repositories
}

func (g *GitHub) ReposByForks(
	ctx context.Context,
) ([]*gh.Repository, error) {
	repositories, err := g.FetchRepos(ctx)
	if err != nil {
		return nil, err
	}

	result := SortByForks(repositories)

	return result, nil
}

func SortByForks(repositories []*gh.Repository) []*gh.Repository {
	// Sort by total forks
	sort.Slice(repositories, func(i, j int) bool {
		return repositories[i].GetForksCount() > repositories[j].GetForksCount()
	})

	return repositories
}

func (g *GitHub) ReposByPRs(
	ctx context.Context,
) ([]RepoAndPRCount, error) {
	repositories, err := g.FetchRepos(ctx)
	if err != nil {
		return nil, err
	}
	// Prefetch PR count before sorting
	repoPRs, err := g.CollectRepoPRCount(ctx, repositories)
	if err != nil {
		return nil, err
	}

	result := SortByPRs(repoPRs)

	return result, nil
}

func SortByPRs(repoPRs []RepoAndPRCount) []RepoAndPRCount {
	// Sort by total PRs
	sort.Slice(repoPRs, func(i, j int) bool {
		return repoPRs[i].PrCount > repoPRs[j].PrCount
	})

	return repoPRs
}

func (g *GitHub) ReposByContributions(
	ctx context.Context,
) ([]RepoAndPRCount, error) {
	repositories, err := g.FetchRepos(ctx)
	if err != nil {
		return nil, err
	}
	// Prefetch PR count before sorting
	repoPRs, err := g.CollectRepoPRCount(ctx, repositories)
	if err != nil {
		return nil, err
	}

	result := SortByContributions(repoPRs)

	return result, nil
}

func SortByContributions(repoPRs []RepoAndPRCount) []RepoAndPRCount {
	// Sort by total PRs
	sort.Slice(repoPRs, func(i, j int) bool {
		return repoPRs[i].PrCount+repoPRs[i].Repo.GetForksCount() > repoPRs[j].PrCount+repoPRs[j].Repo.GetForksCount()
	})

	return repoPRs
}
