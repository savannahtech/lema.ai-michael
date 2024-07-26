package server

import (
	"github.com/dilly3/houdini/internal/model"
	"github.com/dilly3/houdini/internal/repository"
	"github.com/dilly3/houdini/internal/server/mocks"
	"github.com/dilly3/houdini/internal/server/response"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"testing"
)

func TestGetRepoByLanguage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockIRepository(ctrl)
	repository.NewStore(m)

	goRepos := []model.RepoInfo{
		{
			ID:          1,
			Name:        "repo1",
			CreatedAt:   "2021-01-01T00:00:00Z",
			UpdatedAt:   "2021-01-01T00:00:00Z",
			URL:         "test-url.com",
			Description: "test description",
			Language:    "Go",
			Forks:       10,
			Stars:       10,
			OpenIssues:  10,
		},
	}
	m.EXPECT().GetReposByLanguage(gomock.Any(), "Go", 1).Return(goRepos, nil).AnyTimes()
	testCases := []struct {
		name       string
		pattern    string
		path       string
		assertions func(res *response.HttpResponse, code int, err error)
	}{
		{
			name:    "get repos by language success",
			pattern: "/v1/repos/{language}/{limit}",
			path:    "/v1/repos/Go/1",
			assertions: func(res *response.HttpResponse, code int, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, code)
				assert.Equal(t, "repos retrieved successfully", res.Message)
				data := res.Data.([]interface{})
				assert.Equal(t, 1, len(data))
				repo := data[0].(map[string]interface{})
				assert.Equal(t, float64(1), repo["id"])
				assert.Equal(t, "repo1", repo["name"])
				assert.Equal(t, "2021-01-01T00:00:00Z", repo["created_at"])
				assert.Equal(t, "2021-01-01T00:00:00Z", repo["updated_at"])
				assert.Equal(t, "test description", repo["description"])
				assert.Equal(t, "Go", repo["language"])
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, code, err := executeGetRequest(tc.pattern, tc.path, handler.GetReposByLanguage)
			tc.assertions(res, code, err)
		})

	}

}
