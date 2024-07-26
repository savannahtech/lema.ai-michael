package github

import (
	"github.com/mitchellh/mapstructure"
)

func (gh *GHClient) ListCommits(owner, repo string, since string, perPage, page int) ([]CommitResponse, error) {
	var commits []interface{}
	err := gh.listCommits(owner, repo, since, perPage, page, &commits)
	if err != nil {
		return nil, err
	}
	var commitsSlice []CommitResponse
	if len(commits) < 1 {
		return commitsSlice, nil
	}
	for i := 0; i < len(commits); i++ {
		commit := CommitResponse{}
		err = mapstructure.Decode(commits[i], &commit)
		if err != nil {
			return nil, err
		}
		commitsSlice = append(commitsSlice, commit)
	}
	return commitsSlice, nil
}
