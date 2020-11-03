package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/andygrunwald/go-jira"
	"github.com/google/go-github/v32/github"
)

func initJiraClient() {
	tp := jira.BasicAuthTransport{
		Username: settings.Jira.Username,
		Password: settings.Jira.Password,
	}

	client, err := jira.NewClient(tp.Client(), settings.Jira.BaseURL)
	if err != nil {
		panic(fmt.Errorf(`initJiraClient: %s`, err.Error()))
	}

	jiraClient = client
}

func getIssue(issues []string, repoName, head string) ([]jira.Issue, string) {
	url := fmt.Sprintf(`%sbrowse`, settings.Jira.BaseURL)

	releaseNotes := fmt.Sprintf(`Releasenotes %s - %s`, repoName, head) + "\n"

	result := make([]jira.Issue, len(issues))
	for _, i := range issues {
		issue, _, err := jiraClient.Issue.Get(i, nil)
		if err != nil {
			panic(fmt.Errorf(`getIssue: %s`, err.Error()))
		}

		result = append(result, *issue)
		releaseNotes = releaseNotes + "\n" + fmt.Sprintf(`<a href="%s/%s">[%s]</a> - %s - %s`, url, issue.Key, issue.Key, issue.Fields.Type.Name, issue.Fields.Summary)
	}

	return result, releaseNotes
}

func getJiraIdentifiers(commits []*github.RepositoryCommit) []string {
	r, _ := regexp.Compile(fmt.Sprintf("%s-[0-9]{1,5}", settings.Jira.TicketIdentifier))

	tickets := make([]string, 0, len(commits))
	for _, commit := range commits {
		msg := commit.GetCommit().GetMessage()
		if strings.Contains(commit.GetCommit().GetMessage(), settings.Jira.TicketIdentifier) {
			ticket := r.FindString(msg)
			tickets = append(tickets, ticket)
		}
	}

	return tickets
}
