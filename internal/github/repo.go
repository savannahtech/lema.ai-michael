package github

import (
	"context"
	errs "github.com/dilly3/houdini/internal/error"
	"github.com/dilly3/houdini/internal/model"
	"github.com/dilly3/houdini/internal/repository"
	"github.com/dilly3/houdini/internal/repository/cache"
)

func (g *GHubITR) GetRepo(owner, repo string) (*model.RepoInfo, error) {
	res, err := g.ghc.GetRepo(owner, repo)
	if err != nil {
		return nil, errs.NewAppError("failed to get repo", err)
	}
	repoData := mapRepoResponse(res)
	return &repoData, nil
}
func (g *GHubITR) GetRepoCron() error {
	cac := cache.GetDefaultCache()
	res, err := g.ghc.GetRepo(cac.GetOwner(), cac.GetRepo())
	if err != nil {
		return err
	}
	var repoData model.RepoInfo
	repoData = mapRepoResponse(res)
	err = repository.GetDefaultStore().SaveRepo(context.Background(), &repoData)
	if err != nil {
		return err
	}
	return nil
}
