package repository

import (
	"context"
	"github.com/dilly3/houdini/internal/model"
)

// ICommitRepository interface
type ICommitRepository interface {
	GetCommitsByRepoName(ctx context.Context, repoName string, limit int) ([]model.CommitInfo, error)
	GetCommitByID(ctx context.Context, id string) (*model.CommitInfo, error)
	SaveCommit(ctx context.Context, commit *model.CommitInfo) error
	SaveCommits(ctx context.Context, commits []model.CommitInfo) error
	GetLastCommit(ctx context.Context, repoName string) (*model.CommitInfo, error)
	DeleteByDate(ctx context.Context, repoName, date string) error
	GetTopCommitsAuthorsByCount(ctx context.Context, repoName string, limit int) ([]model.AuthorCommits, error)
}
