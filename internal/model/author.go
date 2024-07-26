package model

type AuthorInfo struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty" gorm:"primary_key"`
}

// TableName returns the table name for the AuthorInfo struct
func (AuthorInfo) TableName() string {
	return "authors"
}

// AuthorCommits struct to hold authorCommits
type AuthorCommits struct {
	Author       string `gorm:"author" json:"author"`
	CommitsCount int    `gorm:"commit_count" json:"commit_count"`
}
