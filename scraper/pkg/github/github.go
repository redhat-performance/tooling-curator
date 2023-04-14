package github

import (
	"context"
	"log"

	"github.com/google/go-github/v51/github"
)

func GitHubRepositories(ctx context.Context, org string) []*github.Repository {
	//Using Unauthenticated client
	client := github.NewClient(nil)
	ghrepos, _, err := client.Repositories.List(ctx, org, nil)
	if err != nil {
		log.Fatalf("Error getting repositories from GH: %s", err)
	}
	return ghrepos
}
