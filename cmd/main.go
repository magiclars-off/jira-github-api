package main

import (
	"flag"
	"fmt"
	"jira-api/internal/config"

	"github.com/andygrunwald/go-jira"
	"github.com/google/go-github/v32/github"
)

// Globals
var (
	configFile string
	base       string
	head       string
	repoName   string

	jiraClient *jira.Client
	gitClient  *github.Client
	settings   = struct {
		Jira config.Jira
		Git  config.Git
	}{}
)

func init() {
	flag.StringVar(&configFile, `c`, `config.toml`, `Location of the config file`)
	flag.StringVar(&base, `b`, ``, `Github Base (e.g. tag/v2)`)
	flag.StringVar(&head, `h`, `master`, `Github Head`)
	flag.StringVar(&repoName, `r`, ``, `Repository name`)
	flag.Parse()

	if repoName == "" {
		panic(`specify a repository with (-r)`)
	}
	if base == "" {
		panic(`specify a base (-b) value for comparrison`)
	}
}

func main() {
	defer recoverFunc()

	config.Read(configFile, &settings)

	initJiraClient()
	initGitClient()

	printReleaseNotes()
}

func printReleaseNotes() {
	commits := compareBranches(repoName, base, head)
	ticketIdentifiers := getJiraIdentifiers(commits)
	_, releaseNotes := getIssue(ticketIdentifiers, repoName, head)

	fmt.Print(releaseNotes)
}
