package main

import (
	"testing"

	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
)

func TestGetJiraIdentifiers(t *testing.T) {
	assert := assert.New(t)

	msgWith := "DD-50 - Fix bug"
	input := []*github.RepositoryCommit{
		{
			Commit: &github.Commit{
				Message: &msgWith,
			},
		},
	}

	settings.Jira.TicketIdentifier = "DD"
	t.Run("Valid - identifier found", func(t *testing.T) {
		results := getJiraIdentifiers(input)
		assert.Len(results, 1)
		assert.Equal("DD-50", results[0])
	})

	settings.Jira.TicketIdentifier = "AB"
	t.Run("Valid - nothing found", func(t *testing.T) {
		results := getJiraIdentifiers(input)
		assert.Empty(results)
	})
}
