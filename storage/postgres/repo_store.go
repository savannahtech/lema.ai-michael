package postgres

import (
	"context"
	"github.com/dilly3/houdini/internal/model"
)

// RepoStore struct
type RepoStore struct {
	storage *Storage
}

// NewRepoStore creates a new RepoStore
func NewRepoStore(storage *Storage) *RepoStore {
	return &RepoStore{storage: storage}
}

// SaveRepo saves repo
func (rs *RepoStore) SaveRepo(ctx context.Context, repo *model.RepoInfo) error {
	return rs.storage.DB.WithContext(ctx).FirstOrCreate(repo).Error
}

// GetRepoByID gets repo by id
func (rs *RepoStore) GetRepoByID(ctx context.Context, id string) (*model.RepoInfo, error) {
	var repo model.RepoInfo
	err := rs.storage.DB.WithContext(ctx).Where("id = ?", id).First(&repo).Error
	return &repo, err
}

// GetRepoByName gets repo by name
func (rs *RepoStore) GetRepoByName(ctx context.Context, name string) (*model.RepoInfo, error) {
	var repo model.RepoInfo
	err := rs.storage.DB.WithContext(ctx).Where("LOWER(name) = LOWER(?)", name).First(&repo).Error
	return &repo, err
}

// GetReposByLanguage gets repos by language
func (rs *RepoStore) GetReposByLanguage(ctx context.Context, language string, limit int) ([]model.RepoInfo, error) {
	var repos []model.RepoInfo
	err := rs.storage.DB.WithContext(ctx).Where("LOWER(language) = LOWER(?)", language).Limit(limit).Find(&repos).Error
	return repos, err
}

// GetReposByStarCount gets repos by star count
func (rs *RepoStore) GetReposByStarCount(ctx context.Context, limit int) ([]model.RepoInfo, error) {
	var repos []model.RepoInfo
	err := rs.storage.DB.WithContext(ctx).Order("stars desc").Limit(limit).Find(&repos).Error
	return repos, err
}
