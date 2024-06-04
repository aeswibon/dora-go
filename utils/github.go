package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/aeswibon/dora-go/config"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Github is an interface for interacting with the Github API
type Github interface {
	GetReleaseInfo(owner, repo string) ([]*github.RepositoryRelease, error)
	GetPullRequestInfo(owner, repo string) ([]*github.PullRequest, error)
	GetIssueInfo(owner, repo string) ([]*github.Issue, error)
}

// GithubClient is a struct that implements the Github interface
type GithubClient struct {
	Client *github.Client
}

// NewGithubClient creates a new GithubClient
func NewGithubClient() Github {
	token := config.GetEnv("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)
	return &GithubClient{Client: client}
}

// GetReleaseInfo returns the release information for a given repository
func (g *GithubClient) GetReleaseInfo(owner, repo string) ([]*github.RepositoryRelease, error) {
	releases, _, err := g.Client.Repositories.ListReleases(context.Background(), owner, repo, nil)
	if err != nil {
		return nil, fmt.Errorf("Error fetching releases: %v", err)
	}
	log.Printf("Found %d releases for %s/%s", len(releases), owner, repo)
	return releases, nil
}

// GetPullRequestInfo returns the pull request information for a given repository
func (g *GithubClient) GetPullRequestInfo(owner, repo string) ([]*github.PullRequest, error) {
	prs, _, err := g.Client.PullRequests.List(context.Background(), owner, repo, &github.PullRequestListOptions{
		State: "closed",
	})
	if err != nil {
		return nil, fmt.Errorf("Error fetching pull requests: %v", err)
	}
	log.Printf("Found %d pull requests for %s/%s", len(prs), owner, repo)
	return prs, nil
}

// GetIssueInfo returns the issue information for a given repository
func (g *GithubClient) GetIssueInfo(owner, repo string) ([]*github.Issue, error) {
	issues, _, err := g.Client.Issues.ListByRepo(context.Background(), owner, repo, &github.IssueListByRepoOptions{
		State: "closed",
	})
	if err != nil {
		return nil, fmt.Errorf("Error fetching issues: %v", err)
	}
	log.Printf("Found %d issues for %s/%s", len(issues), owner, repo)
	return issues, nil
}