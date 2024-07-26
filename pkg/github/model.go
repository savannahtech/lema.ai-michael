package github

type CommitResponse struct {
	CommitDetails `json:"commit" mapstructure:"commit"`
}
type CommitDetails struct {
	Message string `json:"message" mapstructure:"message"`
	Author  Author `json:"author" mapstructure:"author"`
	URL     string `json:"url" mapstructure:"url"`
}
type Author struct {
	Name  string `json:"name" mapstructure:"name"`
	Email string `json:"email" mapstructure:"email"`
	Date  string `json:"date" mapstructure:"date"`
}

type RepoResponse struct {
	ID          int    `json:"id" mapstructure:"id"`
	Name        string `json:"name" mapstructure:"name"`
	CreatedAt   string `json:"created_at" mapstructure:"created_at"`
	UpdatedAt   string `json:"updated_at" mapstructure:"updated_at"`
	URL         string `json:"html_url" mapstructure:"html_url"`
	Description string `json:"description" mapstructure:"description"`
	Language    string `json:"language" mapstructure:"language"`
	Forks       int    `json:"forks" mapstructure:"forks"`
	Stars       int    `json:"stargazers_count" mapstructure:"stargazers_count"`
	OpenIssues  int    `json:"open_issues" mapstructure:"open_issues"`
}
