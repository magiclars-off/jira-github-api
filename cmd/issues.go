package main

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
)

func getIssue(issues []string) []jira.Issue {
	url := fmt.Sprintf(`%sbrowse`, settings.Jira.BaseURL)
	
	result := make([]jira.Issue, len(issues))
	for _, i := range issues {
		issue, _, err := jiraClient.Issue.Get(i, nil)
		panicOnError(err)

		result = append(result, *issue)
		fmt.Printf(`<a href="%s/%s">[%s]</a> - %s`, url, issue.Key, issue.Key, issue.Fields.Summary)
		fmt.Printf("\nType: %s\n\n", issue.Fields.Type.Name)
	}

	return result
}
