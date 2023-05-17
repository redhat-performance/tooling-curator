package github

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v51/github"
	"golang.org/x/oauth2"
)

var Contributors []*github.Contributor
var Commits []*github.RepositoryCommit

func GitHubAuth(ctx context.Context) *github.Client {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Token not detected, please export your oauth token as GITHUB_AUTH_TOKEN on the environment")
	}
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}

// Get repos in the org handling pagination
func GitHubRepositories(ctx context.Context, org string, client *github.Client) []*github.Repository {
	opts := &github.RepositoryListOptions{}
	var allrepos []*github.Repository
	for {
		ghrepos, resp, err := client.Repositories.List(ctx, org, opts)
		if err != nil {
			log.Fatalf("Error getting repositories from GH: %s", err)
		}
		allrepos = append(allrepos, ghrepos...)
		if resp.NextPage == 0 {
			break
		}
		opts.ListOptions.Page = resp.NextPage
	}
	return allrepos
}

func ListContrib(ctx context.Context, org string, repository string, client *github.Client) {
	opts := &github.ListContributorsOptions{Anon: "false"}
	var err error
	Contributors, _, err = client.Repositories.ListContributors(ctx, org, repository, opts)
	if err != nil {
		log.Fatalf("Error getting contributors: %s", err)
	}
}

func ListCommits(ctx context.Context, org string, repository string, client *github.Client, lookBack int) {
	opts := &github.CommitsListOptions{Since: time.Now().UTC().AddDate(-lookBack, 0, 0)}
	// Handle case when repo is initialized but there are no commits, by always returning commits
	Commits, _, _ = client.Repositories.ListCommits(ctx, org, repository, opts)
}
