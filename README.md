# GitHubs Top N

A command-line tool that interacts with a given GitHub organization to display
different top-n facts about repositories such as the top 10 most starred repositories.

## Installation

## Usage

```
Usage:
  github-top [command]

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

## How to test

## Note for reviewer
