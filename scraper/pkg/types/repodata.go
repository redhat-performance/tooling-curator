package types

type Repo struct {
	Org         string    `json:"org"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Labels      []string  `json:"labels"`
        Contacts    []Contact `json:"contacts"`
        Maintained  bool      `json:"maintained"`
}

type RepoData struct {
	Repos []Repo `json:"repos"`
}

type Contact struct {
        Username    string   `json:"username"`
        URL         string   `json:"htmlurl"`
}

