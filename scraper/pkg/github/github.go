package github

import (
	"context"
	"log"

	"github.com/google/go-github/github"
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

func ListContrib(ctx context.Context, org string, repository string)  []*github.Contributor {
        //Using Unauthenticated client
        client := github.NewClient(nil)
        opts := &github.ListContributorsOptions{Anon: "false"}
        contributors, _, err :=  client.Repositories.ListContributors(ctx, org, repository, opts) 
        if err != nil {
                log.Fatalf("Error getting contributors: %s", err)
        }
        return contributors
}
