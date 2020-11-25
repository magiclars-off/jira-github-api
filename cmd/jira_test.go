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
		results := getIdentifiersFromCommits(input)
		assert.Len(results, 1)
		assert.Equal("DD-50", results[0])
	})

	settings.Jira.TicketIdentifier = "AB"
	t.Run("Valid - nothing found", func(t *testing.T) {
		results := getIdentifiersFromCommits(input)
		assert.Empty(results)
	})
}

func TestAddCommentToJiraIssue(t *testing.T) {
	assert := assert.New(t)

	t.Run("Panics - no PR template used", func(t *testing.T) {
		assert.PanicsWithValue("seems like the PR didn't use the PR template", func() { determineJiraMessage(`please review this pullRequest`) })
	})

	t.Run("Valid - no PR template used", func(t *testing.T) {
		msg := `### Needs testing from Nestorsupport?
	
		- [x] No
		- [ ] Yes
		
		If Yes, please describe how this change can be tested: 
		
		1. 
		2. 
		3. 
		4.`

		result := determineJiraMessage(msg)
		assert.Equal("No testing of support required", result)
	})

	t.Run("Valid - no PR template used", func(t *testing.T) {
		msg := `### Needs testing from Nestorsupport?
	
		- [ ] No
		- [x] Yes
		
		If Yes, please describe how this change can be tested

		Steps to repoduce:
		1. step 1
		2. step 2
		3. step 3
		4.`

		result := determineJiraMessage(msg)
		assert.Equal(`steps to repoduce:
		1. step 1
		2. step 2
		3. step 3
		4.`, result)
	})

}
