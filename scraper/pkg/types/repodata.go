package types

type Repo struct {
	Org         string   `json:"org"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Labels      []string `json:"labels"`
}

type RepoData struct {
	Repos []Repo `json:"repos"`
}
