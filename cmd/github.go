package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

var ctx = context.Background()

func initGitClient() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: settings.Git.Token},
	)
	tc := oauth2.NewClient(ctx, ts)
	gitClient = github.NewClient(tc)
}

func createRelease(repoName string, release github.RepositoryRelease) int {
	_, resp, err := gitClient.Repositories.CreateRelease(ctx, settings.Git.Owner, repoName, &release)
	if err != nil {
		panic(fmt.Errorf(`createRelease: %s`, err.Error()))
	}

	return resp.StatusCode
}

func compareBranches(repoName, base, head string) []*github.RepositoryCommit {
	comparison, _, err := gitClient.Repositories.CompareCommits(ctx, settings.Git.Owner, repoName, base, head)
	if err != nil {
		panic(fmt.Errorf(`compareBranches: %s`, err.Error()))
	}

	return comparison.Commits
}
