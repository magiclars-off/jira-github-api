package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

var ctx = context.Background()

func (g gitData) initGitClient() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: settings.Git.Token},
	)
	tc := oauth2.NewClient(ctx, ts)
	g.gitClient = github.NewClient(tc)
}

func (g gitData) createRelease(repoName string, release github.RepositoryRelease) int {
	_, resp, err := g.gitClient.Repositories.CreateRelease(ctx, settings.Git.Owner, g.repoName, &release)
	if err != nil {
		panic(fmt.Errorf(`createRelease: %s`, err.Error()))
	}

	return resp.StatusCode
}

func (g gitData) compareBranches() []*github.RepositoryCommit {
	comparison, _, err := g.gitClient.Repositories.CompareCommits(ctx, settings.Git.Owner, g.repoName, g.base, g.head)
	if err != nil {
		panic(fmt.Errorf(`compareBranches: %s`, err.Error()))
	}

	return comparison.Commits
}

func (g gitData) getPullRequest() *github.PullRequest {
	pr, resp, err := g.gitClient.PullRequests.Get(ctx, settings.Git.Owner, g.repoName, 309) //prNr)
	if err != nil {
		panic(fmt.Errorf(`getPullRequest: %s`, err.Error()))
	}
	if resp.StatusCode == 404 {
		panic(fmt.Errorf(`pullRequest nr: %d not found`, g.prNr))
	}

	return pr
}
