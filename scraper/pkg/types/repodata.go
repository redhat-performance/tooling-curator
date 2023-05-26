package types

type Repo struct {
	Org         string    `json:"org"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Labels      []string  `json:"labels"`
	Contacts    []Contact `json:"contacts"`
	Active      bool      `json:"active"`
	Archived    bool      `json:"archived"`
}

type RepoData struct {
	Repos []Repo `json:"repos"`
}

type Contact struct {
	Username string `json:"username"`
	URL      string `json:"htmlurl"`
}

// ignored-repositories.json struct
type IgnoredRepositories struct {
	Global         []string              `json:"global"`
	Orgs           map[string]OrgsConfig `json:"orgs"`
	IgnoreArchived bool                  `json:"ignoreArchived"`
}

type OrgsConfig struct {
	SkipGlobalIgnore   bool     `json:"skip-global-ignore"`
	SkipGlobalArchived bool     `json:"skip-global-archived"`
	Repos              []string `json:"repos"`
}
