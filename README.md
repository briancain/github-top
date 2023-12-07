# GitHub's Top Metrics

A command-line tool that interacts with a GitHub organization to display
different top facts about repositories such as the top 10 most starred repositories.

## Installation

To install the latest release of this tool, head over to the [releases](https://github.com/briancain/github-top/releases)
page on GitHub! Pick your current platform (i.e. if on macOS with an M1, choose
darwin_arm64). Unpackage the archive and place the binary on your path or
run it directly:

```shell
$ tar xvf github-top_0.0.1_darwin_arm64.tar.gz
x github-top
$ ./github-top --help
...
```

__NOTE__: I did not have time to sign this properly with a developer key to run
on MacOS. If you are on macOS, you might need to build the tool yourself.

Clone the repository and build with Go:

```shell
$ git clone git@github.com:briancain/github-top.git
$ cd github-top
$ make
...
$ ./bin/github-top --help
...
```

## Usage

```
A command line tool to interact with GitHub organizations

Usage:
  github-top [command]

Examples:
github-top stars --org-name netflix

Available Commands:
  completion    Generate the autocompletion script for the specified shell
  contributions Display the top repos in an org by total contributions (PRs and forks)
  forks         Display the top repos in an org by total forks
  help          Help about any command
  pull-requests Display the top repos in an org by total pull requests
  stars         Display the top repos in an org by total stars

Flags:
  -h, --help              help for github-top
  -o, --org-name string   The GitHub organization to query
      --token string      The GitHub token to use
  -n, --total int         The total results to display (default 10)

Use "github-top [command] --help" for more information about a command.
```

GitHub allows for a few unathenticated API requests per-day. You will likely
run over that limit after a few requests. Once that happens, you will need
to [generate an API Key](https://github.com/settings/tokens) to use with the `--token` flag.

## How to test

Pick your favorite organization on GitHub and figure out the top statistics:

Some good Organizations to test:

- netflix
- hashicorp
- spinnaker
- cookiecutter
    + This org is small and only has a few repos

```shell
$ ./bin/github-top --token <TOKEN_HERE> stars -o netflix
```

## Note for the user

This CLI tool uses that API via [go-github](https://github.com/google/go-github). For large queries
with organizations with many repositories, this command line tool will iterate
through each paginated list of repos (100 per-page) to collect all repos for an org.

Because of this, during my testing I noticed the GitHub API is quite slow for large queries
especially when calculating Pull requests for each repo. Both collecting all
repositories and all pull requests per-repo require iterating through each page,
making the CLI tools rather slow for large queries using the REST API.
Ideally, an optimization here would be to use GitHub's [GraphQL](https://docs.github.com/en/graphql)
API. Then this CLI could limit the results based on the request from the user
rather than querying for the entire data set, iterating over each page, and
calculating the top result itself.
