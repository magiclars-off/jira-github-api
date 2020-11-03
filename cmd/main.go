package main

import (
	"flag"
	"fmt"
	"jira-api/internal/config"
	"os"

	"github.com/andygrunwald/go-jira"
	"github.com/google/go-github/v32/github"
)

// Globals
var (
	jiraClient *jira.Client
	gitClient  *github.Client
	settings   = struct {
		Jira config.Jira
		Git  config.Git
	}{}
)

func main() {
	defer recoverFunc()

	configFile := flag.String(`c`, `config.toml`, `Location of the config file`)
	base := flag.String(`b`, ``, `Github Base (e.g. tag/v2)`)
	head := flag.String(`h`, `master`, `Github Head`)
	repoName := flag.String(`r`, ``, `Repository name`)

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage of Jira-api tool:")
		flag.PrintDefaults()

		fmt.Println("\nMake sure to also set env variables or config file")
	}
	flag.Parse()

	if *repoName == "" {
		panic(`specify a repository with (-r)`)
	}
	if *base == "" || *head == "" {
		panic(`specify a base (-b) and head (-h) values for comparrison`)
	}

	config.Read(*configFile, &settings)

	initJiraClient()
	initGitClient()

	commits := compareBranches(*repoName, *base, *head)
	ticketIdentifiers := getJiraIdentifiers(commits)
	_, releaseNotes := getIssue(ticketIdentifiers, *repoName, *head)

	fmt.Print(releaseNotes)
}
