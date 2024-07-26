package model

type CommitInfo struct {
	ID          string `gorm:"primary_key"`
	RepoName    string `json:"repo_name"`
	Message     string `gorm:"index"`
	AuthorName  string
	AuthorEmail string
	Date        string
	URL         string
}

// TableName returns the table name for the CommitInfo struct
func (CommitInfo) TableName() string {
	return "commits"
}
