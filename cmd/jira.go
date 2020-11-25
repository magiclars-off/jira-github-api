package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/andygrunwald/go-jira"
	"github.com/google/go-github/v32/github"
)

func (j jiraData) initJiraClient() {
	tp := jira.BasicAuthTransport{
		Username: settings.Jira.Username,
		Password: settings.Jira.Password,
	}

	client, err := jira.NewClient(tp.Client(), settings.Jira.BaseURL)
	if err != nil {
		panic(fmt.Errorf(`initJiraClient: %s`, err.Error()))
	}

	j.jiraClient = client
}

func (j jiraData) getJiraIssue(id string) *jira.Issue {
	issue, resp, err := j.jiraClient.Issue.Get(id, nil)
	if err != nil {
		panic(fmt.Errorf(`getJiraIssue: %s`, err.Error()))
	}
	if resp.StatusCode == 404 {
		panic(fmt.Errorf(`jiraIssue: %s not found`, id))
	}

	return issue
}

func (j jiraData) addCommentToJiraIssue(id string, message string) {
	_, resp, err := j.jiraClient.Issue.AddComment(id, &jira.Comment{
		Body: message,
	})
	if err != nil {
		panic(fmt.Errorf(`addCommentToJiraIssue: %s`, err.Error()))
	}
	if resp.StatusCode != 200 {
		panic(fmt.Errorf(`addCommentToJiraIssue failed with statusCode: %d`, resp.StatusCode))
	}
}

func getIdentifiersFromCommits(commits []*github.RepositoryCommit) []string {
	tickets := make([]string, 0, len(commits))
	for _, commit := range commits {
		if ticket := findJiraIdentifier(commit.GetCommit().GetMessage()); ticket != "" {
			tickets = append(tickets, ticket)
		}
	}

	return tickets
}

func findJiraIdentifier(message string) string {
	r, _ := regexp.Compile(fmt.Sprintf("%s-[0-9]{1,5}", settings.Jira.TicketIdentifier))
	return r.FindString(message)
}

func determineJiraMessage(message string) string {
	start := "Needs testing from Nestorsupport?"
	index := strings.Index(message, start)
	if index == -1 {
		panic("seems like the PR didn't use the PR template")
	}

	if strings.Contains(message, "- [x] No") {
		return "No testing of support required"
	}

	shortMessage := strings.Split(message, start)
	fmt.Print(shortMessage[1])
	return message
}
