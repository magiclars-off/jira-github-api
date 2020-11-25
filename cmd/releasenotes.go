package main

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
)

func (g gitData) validate() {
	if g.repoName == "" {
		panic(`specify a repository with (-r)`)
	}
	if g.base == "" {
		panic(`specify a base (-b) value for comparrison`)
	}
}

func (g gitData) getReleasenotes(issues []jira.Issue) string {
	url := fmt.Sprintf(`%sbrowse`, settings.Jira.BaseURL)
	releaseNotes := fmt.Sprintf(`Releasenotes %s - %s`, g.repoName, g.head) + "\n"

	for _, issue := range issues {
		releaseNotes = releaseNotes + "\n" + fmt.Sprintf(`<a href="%s/%s">[%s]</a> - %s - %s`, url, issue.Key, issue.Key, issue.Fields.Type.Name, issue.Fields.Summary)
	}

	return releaseNotes
}
