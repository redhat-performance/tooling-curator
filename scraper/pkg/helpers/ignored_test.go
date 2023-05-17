package helpers

import (
	"testing"

	gh "github.com/google/go-github/v51/github"
	"github.com/redhat-performance/tooling-curator/scraper/pkg/types"
)

// Repositories definitions
var (
	testRepo = gh.Repository{
		Archived: new(bool),
		Name:     new(string),
		Owner:    new(gh.User),
	}

	dotGithubRepo = gh.Repository{
		Archived: new(bool),
		Name:     new(string),
		Owner:    new(gh.User),
	}

	archivedRepo = gh.Repository{
		Archived: new(bool),
		Name:     new(string),
		Owner:    new(gh.User),
	}

	otherOrgRepo = gh.Repository{
		Archived: new(bool),
		Name:     new(string),
		Owner:    new(gh.User),
	}

	topicIgnoredRepo = gh.Repository{
		Archived: new(bool),
		Name:     new(string),
		Owner:    new(gh.User),
	}

	archivedRepoLocalIgnored = gh.Repository{
		Archived: new(bool),
		Name:     new(string),
		Owner:    new(gh.User),
	}
)

// Rules definitions
var (
	global_dot_github_ignore = types.IgnoredRepositories{
		Global:         []string{".github"},
		IgnoreArchived: false,
	}
	global_dot_github_ignore_archive_ignore = types.IgnoredRepositories{
		Global:         []string{".github"},
		IgnoreArchived: true,
	}
	org_specific_ignore = types.IgnoredRepositories{
		Global:         []string{},
		IgnoreArchived: false,
		Orgs:           map[string]types.OrgsConfig{},
	}
	global_dot_github_ignore_archive_ignore_org_skip_global_ignore = types.IgnoredRepositories{
		Global:         []string{".github"},
		IgnoreArchived: true,
		Orgs:           map[string]types.OrgsConfig{},
	}
	global_dot_github_ignore_archive_ignore_org_skip_global_archive = types.IgnoredRepositories{
		Global:         []string{".github"},
		IgnoreArchived: true,
		Orgs:           map[string]types.OrgsConfig{},
	}
	global_dot_github_ignore_archive_ignore_org_skip_global_archive_specific_ignore = types.IgnoredRepositories{
		Global:         []string{".github"},
		IgnoreArchived: true,
		Orgs:           map[string]types.OrgsConfig{},
	}
)

var ignoreTopics = []string{"golang"}

func initTestData() {
	// Initializing testRepo
	*testRepo.Archived = false
	*testRepo.Name = "test-repo"
	testRepo.Owner.Login = new(string)
	*testRepo.Owner.Login = "cloud-bulldozer"
	testRepo.Topics = []string{"bash", "python", "not-ignored"}

	// Initializing .github repo
	*dotGithubRepo.Archived = false
	*dotGithubRepo.Name = ".github"
	dotGithubRepo.Owner.Login = new(string)
	*dotGithubRepo.Owner.Login = "cloud-bulldozer"
	dotGithubRepo.Topics = []string{"config", "yaml"}

	// Initializing archived repo
	*archivedRepo.Archived = true
	*archivedRepo.Name = "ye-old-repo"
	archivedRepo.Owner.Login = new(string)
	*archivedRepo.Owner.Login = "cloud-bulldozer"
	archivedRepo.Topics = []string{"ruby"}

	// Initializing otherOrgRepo repo
	*otherOrgRepo.Archived = false
	*otherOrgRepo.Name = "ye-old-repo"
	otherOrgRepo.Owner.Login = new(string)
	*otherOrgRepo.Owner.Login = "distributed-system-analysis"
	otherOrgRepo.Topics = []string{"ruby", "analytics"}

	// Initializing topicIgnoredRepo repo
	*topicIgnoredRepo.Archived = false
	*topicIgnoredRepo.Name = "golang-repo"
	topicIgnoredRepo.Owner.Login = new(string)
	*topicIgnoredRepo.Owner.Login = "cloud-bulldozer"
	topicIgnoredRepo.Topics = []string{"golang", "cli"}

	// Initializing archived repo
	*archivedRepoLocalIgnored.Archived = true
	*archivedRepoLocalIgnored.Name = "ye-old-repo-II"
	archivedRepoLocalIgnored.Owner.Login = new(string)
	*archivedRepoLocalIgnored.Owner.Login = "cloud-bulldozer"
	archivedRepoLocalIgnored.Topics = []string{"ruby"}

	// Initializing IgnoreRules
	org_specific_ignore.Orgs["cloud-bulldozer"] = types.OrgsConfig{
		SkipGlobalIgnore:   false,
		SkipGlobalArchived: false,
		Repos:              []string{"test-repo"}}

	global_dot_github_ignore_archive_ignore_org_skip_global_ignore.Orgs["cloud-bulldozer"] = types.OrgsConfig{
		SkipGlobalIgnore:   true,
		SkipGlobalArchived: false,
		Repos:              []string{"test-repo"}}

	global_dot_github_ignore_archive_ignore_org_skip_global_archive.Orgs["cloud-bulldozer"] = types.OrgsConfig{
		SkipGlobalIgnore:   false,
		SkipGlobalArchived: true,
		Repos:              []string{"test-repo"}}

	global_dot_github_ignore_archive_ignore_org_skip_global_archive_specific_ignore.Orgs["cloud-bulldozer"] = types.OrgsConfig{
		SkipGlobalIgnore:   false,
		SkipGlobalArchived: true,
		Repos:              []string{"test-repo", "ye-old-repo-II"}}
}

func TestIgnored(t *testing.T) {
	initTestData()
	type args struct {
		r  gh.Repository
		o  string
		ir types.IgnoredRepositories
		it []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// GlobalIgnore ".github", GlobalArchivedIgnore "false"
		{"testRepo-global_dot_github_ignore", args{testRepo, "cloud-bulldozer", global_dot_github_ignore, ignoreTopics}, false},
		{"dotGithubRepo-via-global-global_dot_github_ignore", args{dotGithubRepo, "cloud-bulldozer", global_dot_github_ignore, ignoreTopics}, true},
		{"archivedRepo-global_dot_github_ignore", args{archivedRepo, "cloud-bulldozer", global_dot_github_ignore, ignoreTopics}, false},
		{"otherOrgRepo-global_dot_github_ignore", args{otherOrgRepo, "distributed-system-analysis", global_dot_github_ignore, ignoreTopics}, false},
		{"topicIgnoredRepo-global_dot_github_ignore", args{topicIgnoredRepo, "cloud-bulldozer", global_dot_github_ignore, ignoreTopics}, true},
		{"archivedRepoLocalIgnored-global_dot_github_ignore", args{archivedRepoLocalIgnored, "cloud-bulldozer", global_dot_github_ignore, ignoreTopics}, false},

		// GlobalIgnore ".github", GlobalArchivedIgnore "true"
		{"testRepo-global_dot_github_ignore_archive_ignore", args{testRepo, "cloud-bulldozer", global_dot_github_ignore_archive_ignore, ignoreTopics}, false},
		{"dotGithubRepo-via-global-global_dot_github_ignore_archive_ignore", args{dotGithubRepo, "cloud-bulldozer", global_dot_github_ignore_archive_ignore, ignoreTopics}, true},
		{"archivedRepo-global_dot_github_ignore_archive_ignore", args{archivedRepo, "cloud-bulldozer", global_dot_github_ignore_archive_ignore, ignoreTopics}, true},
		{"otherOrgRepo-global_dot_github_ignore_archive_ignore", args{otherOrgRepo, "distributed-system-analysis", global_dot_github_ignore_archive_ignore, ignoreTopics}, false},
		{"topicIgnoredRepo-global_dot_github_ignore_archive_ignore", args{topicIgnoredRepo, "cloud-bulldozer", global_dot_github_ignore_archive_ignore, ignoreTopics}, true},
		{"archivedRepoLocalIgnored-global_dot_github_ignore_archive_ignore", args{archivedRepoLocalIgnored, "cloud-bulldozer", global_dot_github_ignore_archive_ignore, ignoreTopics}, true},

		// Org ignore "test-repo"
		{"testRepo-org_specific_ignore", args{testRepo, "cloud-bulldozer", org_specific_ignore, ignoreTopics}, true},
		{"dotGithubRepo-via-global-org_specific_ignore", args{dotGithubRepo, "cloud-bulldozer", org_specific_ignore, ignoreTopics}, false},
		{"archivedRepo-org_specific_ignore", args{archivedRepo, "cloud-bulldozer", org_specific_ignore, ignoreTopics}, false},
		{"otherOrgRepo-org_specific_ignore", args{otherOrgRepo, "distributed-system-analysis", org_specific_ignore, ignoreTopics}, false},
		{"topicIgnoredRepo-org_specific_ignore", args{topicIgnoredRepo, "cloud-bulldozer", org_specific_ignore, ignoreTopics}, true},
		{"archivedRepoLocalIgnored-org_specific_ignore", args{archivedRepoLocalIgnored, "cloud-bulldozer", org_specific_ignore, ignoreTopics}, false},

		// GlobalIgnore ".github", GlobalArchivedIgnore "true", OrgIgnore "test-repo", SkipGlobalIgnore True
		{"testRepo-global_dot_github_ignore_archive_ignore_org_skip_global_ignore", args{testRepo, "cloud-bulldozer", global_dot_github_ignore_archive_ignore_org_skip_global_ignore, ignoreTopics}, true},
		{"dotGithubRepo-via-global-global_dot_github_ignore_archive_ignore_org_skip_global_ignore", args{dotGithubRepo, "cloud-bulldozer", global_dot_github_ignore_archive_ignore_org_skip_global_ignore, ignoreTopics}, false},
		{"archivedRepo-global_dot_github_ignore_archive_ignore_org_skip_global_ignore", args{archivedRepo, "cloud-bulldozer", global_dot_github_ignore_archive_ignore_org_skip_global_ignore, ignoreTopics}, true},
		{"otherOrgRepo-global_dot_github_ignore_archive_ignore_org_skip_global_ignore", args{otherOrgRepo, "distributed-system-analysis", global_dot_github_ignore_archive_ignore_org_skip_global_ignore, ignoreTopics}, false},
		{"topicIgnoredRepo-global_dot_github_ignore_archive_ignore_org_skip_global_ignore", args{topicIgnoredRepo, "cloud-bulldozer", global_dot_github_ignore_archive_ignore_org_skip_global_ignore, ignoreTopics}, true},
		{"archivedRepoLocalIgnored-global_dot_github_ignore_archive_ignore_org_skip_global_ignore", args{archivedRepoLocalIgnored, "cloud-bulldozer", global_dot_github_ignore_archive_ignore_org_skip_global_ignore, ignoreTopics}, true},

		// GlobalIgnore ".github", GlobalArchivedIgnore "true", OrgIgnore "test-repo", SkipGlobalArchived True
		{"testRepo-global_dot_github_ignore_archive_ignore_org_skip_global_archive", args{testRepo, "cloud-bulldozer", global_dot_github_ignore_archive_ignore_org_skip_global_archive, ignoreTopics}, true},
		{"dotGithubRepo-via-global-global_dot_github_ignore_archive_ignore_org_skip_global_archive", args{dotGithubRepo, "cloud-bulldozer", global_dot_github_ignore_archive_ignore_org_skip_global_archive, ignoreTopics}, true},
		{"archivedRepo-global_dot_github_ignore_archive_ignore_org_skip_global_archive", args{archivedRepo, "cloud-bulldozer", global_dot_github_ignore_archive_ignore_org_skip_global_archive, ignoreTopics}, false},
		{"otherOrgRepo-global_dot_github_ignore_archive_ignore_org_skip_global_archive", args{otherOrgRepo, "distributed-system-analysis", global_dot_github_ignore_archive_ignore_org_skip_global_archive, ignoreTopics}, false},
		{"topicIgnoredRepo-global_dot_github_ignore_archive_ignore_org_skip_global_archive", args{topicIgnoredRepo, "cloud-bulldozer", global_dot_github_ignore_archive_ignore_org_skip_global_archive, ignoreTopics}, true},
		{"archivedRepoLocalIgnored-global_dot_github_ignore_archive_ignore_org_skip_global_archive", args{archivedRepoLocalIgnored, "cloud-bulldozer", global_dot_github_ignore_archive_ignore_org_skip_global_archive, ignoreTopics}, false},

		// GlobalIgnore ".github", GlobalArchivedIgnore "true", OrgIgnore "test-repo, ye-old-repo-II", SkipGlobalArchived True
		{"testRepo-global_dot_github_ignore_archive_ignore_org_skip_global_archive_specific_ignore", args{testRepo, "cloud-bulldozer", global_dot_github_ignore_archive_ignore_org_skip_global_archive_specific_ignore, ignoreTopics}, true},
		{"dotGithubRepo-via-global-global_dot_github_ignore_archive_ignore_org_skip_global_archive_specific_ignore", args{dotGithubRepo, "cloud-bulldozer", global_dot_github_ignore_archive_ignore_org_skip_global_archive_specific_ignore, ignoreTopics}, true},
		{"archivedRepo-global_dot_github_ignore_archive_ignore_org_skip_global_archive_specific_ignore", args{archivedRepo, "cloud-bulldozer", global_dot_github_ignore_archive_ignore_org_skip_global_archive_specific_ignore, ignoreTopics}, false},
		{"otherOrgRepo-global_dot_github_ignore_archive_ignore_org_skip_global_archive_specific_ignore", args{otherOrgRepo, "distributed-system-analysis", global_dot_github_ignore_archive_ignore_org_skip_global_archive_specific_ignore, ignoreTopics}, false},
		{"topicIgnoredRepo-global_dot_github_ignore_archive_ignore_org_skip_global_archive_specific_ignore", args{topicIgnoredRepo, "cloud-bulldozer", global_dot_github_ignore_archive_ignore_org_skip_global_archive_specific_ignore, ignoreTopics}, true},
		{"archivedRepoLocalIgnored-global_dot_github_ignore_archive_ignore_org_skip_global_archive_specific_ignore", args{archivedRepoLocalIgnored, "cloud-bulldozer", global_dot_github_ignore_archive_ignore_org_skip_global_archive_specific_ignore, ignoreTopics}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ignored(tt.args.r, tt.args.o, tt.args.ir, tt.args.it); got != tt.want {
				t.Errorf("%v - Ignored() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
