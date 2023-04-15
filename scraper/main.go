package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/redhat-performance/tooling-curator/scraper/pkg/github"
	"github.com/redhat-performance/tooling-curator/scraper/pkg/types"
)

const (
	organizationsFile      = "../public/organizations.json"
	repositoriesFile       = "../public/repositories.json"
	ignoredTopicsFile      = "../public/ignored-topics.json"
	ignoreRepositoriesFile = "../public/ignored-repositories.json"
	topContributorsCount   = 3
)

var (
	ctx                                      = context.Background()
	sOrgs               *[]string            = &[]string{}
	ignoredTopics       *[]string            = &[]string{}
	ignoredRepositories *map[string][]string = &map[string][]string{}
)

func loadConfiguration() {
	//Loading organizations file
	orgs, err := os.ReadFile(organizationsFile)
	if err != nil {
		log.Fatalf("Error Reading organizations: %s", err)
	}
	err = json.Unmarshal(orgs, sOrgs)
	if err != nil {
		log.Fatalf("Error Unmarshaling Organizations: %s", err)
	}

	//Loading IgnoredTopics file
	iTopics, err := os.ReadFile(ignoredTopicsFile)
	if err != nil {
		log.Fatalf("Error Reading ignored Topics: %s", err)
	}
	err = json.Unmarshal(iTopics, ignoredTopics)
	if err != nil {
		log.Fatalf("Error Unmarshaling ignored topics: %s", err)
	}

	//Loading IgnoredTopics file
	iRepos, err := os.ReadFile(ignoreRepositoriesFile)
	if err != nil {
		log.Fatalf("Error Reading ignored Repositories: %s", err)
	}
	err = json.Unmarshal(iRepos, ignoredRepositories)
	if err != nil {
		log.Fatalf("Error Unmarshaling ignored Repositories: %s", err)
	}
}

func contains(value string, items []string) bool {
	for _, i := range items {
		if i == value {
			return true
		}
	}
	return false
}

func main() {
	loadConfiguration()
	client := github.GitHubAuth(ctx)
	var repoData types.RepoData
	ir := *ignoredRepositories
	for _, o := range *sOrgs {
		ghrepos := github.GitHubRepositories(ctx, o, client)

		for _, r := range ghrepos {
			ignored := false
			if iRepos, ok := ir[o]; ok {
				ignored = contains(r.GetName(), iRepos)
			}
			if !ignored {
				for _, v := range *ignoredTopics {
					ignored = contains(v, r.Topics)
					if ignored {
						break
					}
				}
			}
			if !ignored {
				var contactData []types.Contact
				topics := r.Topics
				contributors := github.ListContrib(ctx, r.Owner.GetLogin(), r.GetName(), client)
				for n, contributor := range contributors {
					if n > topContributorsCount-1 {
						break
					}
					contacts := types.Contact{Username: *contributor.Login, URL: *contributor.HTMLURL}
					contactData = append(contactData, contacts)
				}
				repo := types.Repo{Org: r.Owner.GetLogin(), Name: r.GetName(), URL: r.GetHTMLURL(), Description: r.GetDescription(), Labels: topics, Contacts: contactData}
				repoData.Repos = append(repoData.Repos, repo)

			}
		}
	}

	reposJson, err := json.Marshal(repoData)
	if err != nil {
		log.Fatalf("Error marshaling Repositories: %s", err)
	}
	//        fmt.Println(repoData.Repos)
	os.WriteFile(repositoriesFile, reposJson, 0666)
}
