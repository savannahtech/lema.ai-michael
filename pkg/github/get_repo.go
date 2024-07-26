package github

import (
	"github.com/mitchellh/mapstructure"
)

func (gh *GHClient) GetRepo(owner, repo string) (*RepoResponse, error) {
	expectedResponse := map[string]interface{}{}
	err := gh.getRepo(owner, repo, &expectedResponse)
	if err != nil {
		return nil, err
	}
	resultFromRepo := RepoResponse{}
	err = mapstructure.Decode(expectedResponse, &resultFromRepo)
	if err != nil {
		return nil, err
	}

	return &resultFromRepo, nil
}
