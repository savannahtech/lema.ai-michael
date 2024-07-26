package model

type RepoInfo struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Name        string `json:"name" `
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	URL         string `json:"html_url"`
	Description string `json:"description"`
	Language    string `json:"language" `
	Forks       int    `json:"forks"`
	Stars       int    `json:"stargazers_count"`
	OpenIssues  int    `json:"open_issues"`
}

// TableName returns the table name for the RepoInfo struct
func (RepoInfo) TableName() string {
	return "repos"
}
