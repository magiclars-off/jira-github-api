package main

import (
	"context"
	"flag"
	"jira-api/internal/config"
	"strings"

	"github.com/andygrunwald/go-jira"
	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
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

	configFile := flag.String(`config`, `config.toml`, `Location of the config file`)
	issues := flag.String(`issues`, ``, `Get issue information`)
	flag.Parse()

	config.Read(*configFile, &settings)

	initJiraClient()
	initGitClient()

	if issues != nil {
		_ = getIssue(strings.Split(*issues, ","))
	}
}

func initJiraClient() {
	tp := jira.BasicAuthTransport{
		Username: settings.Jira.Username,
		Password: settings.Jira.Password,
	}

	client, err := jira.NewClient(tp.Client(), settings.Jira.BaseURL)
	panicOnError(err)
	jiraClient = client
}

func initGitClient() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: settings.Git.Token},
	)
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	gitClient = github.NewClient(tc)
}
