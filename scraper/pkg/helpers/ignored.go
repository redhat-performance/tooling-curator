package helpers

import (
	gh "github.com/google/go-github/v51/github"
	"github.com/redhat-performance/tooling-curator/scraper/pkg/types"
)

func Ignored(r gh.Repository, o string, ir types.IgnoredRepositories, it []string) bool {
	ignored := false
	iRepo, ok := ir.Orgs[o]
	if !ok {
		iRepo = types.OrgsConfig{}
	}

	if r.GetArchived() {
		if !iRepo.SkipGlobalArchived {
			ignored = ir.IgnoreArchived
		}
	}

	rName := r.GetName()
	if !ignored {
		ignored = Contains(rName, iRepo.Repos)
	}

	if !ignored {
		if !iRepo.SkipGlobalIgnore {
			ignored = Contains(rName, ir.Global)
		}
	}

	if !ignored {
		for _, v := range it {
			ignored = Contains(v, r.Topics)
			if ignored {
				break
			}
		}
	}
	return ignored
}
