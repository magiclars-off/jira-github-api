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
	mode       string
	settings   = struct {
		Jira config.Jira
		Git  config.Git
	}{}

	g gitData
	j jiraData
)

type jiraData struct {
	jiraClient *jira.Client
}

type gitData struct {
	prNr      int
	base      string
	head      string
	repoName  string
	gitClient *github.Client
}

func init() {
	flag.StringVar(&mode, `m`, ``, `'support-instructions' or 'releasenotes'`)
	flag.IntVar(&g.prNr, `pr`, 0, `required for: 'support-instructions', pullRequest ID to copy the instructions from`)
	flag.StringVar(&configFile, `c`, `config.toml`, `Location of the config file`)
	flag.StringVar(&g.base, `b`, ``, `Github Base (e.g. tag/v2)`)
	flag.StringVar(&g.head, `h`, `master`, `Github Head`)
	flag.StringVar(&g.repoName, `r`, ``, `Repository name`)
}

func main() {
	defer recoverFunc()

	flag.Parse()
	config.Read(configFile, &settings)

	j.initJiraClient()
	g.initGitClient()

	switch mode {
	case config.ModusReleasenotes.String():
		releaseNotes()
		break
	case config.ModusSupportInstructions.String():
		supportInstructions()
		break
	default:
		panic(fmt.Errorf(`'%s' is not a valid mode`, mode))
	}
}

func supportInstructions() {
	pr := g.getPullRequest()

	jiraID := findJiraIdentifier(*pr.Title)
	if jiraID == "" {
		panic("no JiraIdentifier found in the PR title")
	}

	issue := j.getJiraIssue(jiraID)
	msg := determineJiraMessage(*pr.Body)
	j.addCommentToJiraIssue(issue.ID, msg)
}

// releaseNotes takes the input branches, repository and user. Compares the branches
// checks the commits for jira identifiers and produces releasenotes
func releaseNotes() {
	g.validate()

	commits := g.compareBranches()
	ticketIdentifiers := getIdentifiersFromCommits(commits)

	issues := make([]jira.Issue, len(ticketIdentifiers))
	for _, id := range ticketIdentifiers {
		issue := j.getJiraIssue(id)
		if issue != nil {
			issues = append(issues, *issue)
		}
	}

	releaseNotes := g.getReleasenotes(issues)
	fmt.Print(releaseNotes)
}
