package github

import (
	"context"
	"log"
        "os"

	"github.com/google/go-github/v51/github"
        "golang.org/x/oauth2"
)

func GitHubAuth(ctx context.Context) *github.Client {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Token not detected, please export your oauth token as GITHUB_AITH_TOKEN on the environment")
	}
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
        client := github.NewClient(tc)
        return client
}

func GitHubRepositories(ctx context.Context, org string, client *github.Client) []*github.Repository {
	ghrepos, _, err := client.Repositories.List(ctx, org, nil)
	if err != nil {
		log.Fatalf("Error getting repositories from GH: %s", err)
	}
	return ghrepos
}

func ListContrib(ctx context.Context, org string, repository string, client *github.Client)  []*github.Contributor {
        opts := &github.ListContributorsOptions{Anon: "false"}
        contributors, _, err :=  client.Repositories.ListContributors(ctx, org, repository, opts) 
        if err != nil {
                log.Fatalf("Error getting contributors: %s", err)
        }
        return contributors
}
