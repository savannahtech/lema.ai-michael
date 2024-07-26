package postgres

import (
	"context"
	"github.com/dilly3/houdini/internal/model"
)

// CommitStore struct
type CommitStore struct {
	storage *Storage
}

// NewCommitStore creates a new CommitStore
func NewCommitStore(storage *Storage) *CommitStore {
	return &CommitStore{storage: storage}
}

// GetCommitsByRepoName gets commits by repo name
func (cs *CommitStore) GetCommitsByRepoName(ctx context.Context, repoName string, limit int) ([]model.CommitInfo, error) {
	var commits []model.CommitInfo
	err := cs.storage.DB.WithContext(ctx).Where("LOWER(repo_name) = LOWER(?)", repoName).Order("date desc").Limit(limit).Find(&commits).Error
	return commits, err
}

// GetCommitByID gets commit by id
func (cs *CommitStore) GetCommitByID(ctx context.Context, id string) (*model.CommitInfo, error) {
	var commit model.CommitInfo
	err := cs.storage.DB.WithContext(ctx).Where("id = ?", id).First(&commit).Error
	return &commit, err
}

// SaveCommit saves commit
func (cs *CommitStore) SaveCommit(ctx context.Context, commit *model.CommitInfo) error {
	return cs.storage.DB.WithContext(ctx).FirstOrCreate(commit).Error
}

// SaveCommits saves commits
func (cs *CommitStore) SaveCommits(ctx context.Context, commits []model.CommitInfo) error {
	return cs.storage.DB.WithContext(ctx).Save(commits).Error
}

// GetLastCommit gets last commit (sort by date)
func (cs *CommitStore) GetLastCommit(ctx context.Context, repoName string) (*model.CommitInfo, error) {
	var commit model.CommitInfo
	err := cs.storage.DB.WithContext(ctx).Where("repo_name = ?", repoName).Order("date desc").First(&commit).Error
	return &commit, err
}

// GetTopCommitsAuthorsByCount gets the top commit authors by count
func (cs *CommitStore) GetTopCommitsAuthorsByCount(ctx context.Context, repoName string, limit int) ([]model.AuthorCommits, error) {
	var authorCommits []model.AuthorCommits
	err := cs.storage.DB.WithContext(ctx).Model(&model.CommitInfo{}).Select("author_name as author, COUNT(*) as commits_count").Where("repo_name = ?",
		repoName).Group("author_name").Order("commits_count desc").Limit(limit).Find(&authorCommits).Error
	if err != nil {
		return nil, err
	}
	return authorCommits, nil
}

// DeleteByDate hard deletes commits by date
func (cs *CommitStore) DeleteByDate(ctx context.Context, repoName, date string) error {
	return cs.storage.DB.WithContext(ctx).Delete(&model.CommitInfo{}, "date > ? AND repo_name = ? ", date, repoName).Error
}
