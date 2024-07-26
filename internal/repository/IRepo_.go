package repository

import (
	"context"
	"github.com/dilly3/houdini/internal/model"
)

// IRepoRepository interface
type IRepoRepository interface {
	SaveRepo(ctx context.Context, repo *model.RepoInfo) error
	GetRepoByID(ctx context.Context, id string) (*model.RepoInfo, error)
	GetRepoByName(ctx context.Context, name string) (*model.RepoInfo, error)
	GetReposByLanguage(ctx context.Context, language string, limit int) ([]model.RepoInfo, error)
	GetReposByStarCount(ctx context.Context, limit int) ([]model.RepoInfo, error)
}
