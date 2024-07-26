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

func TestGetRepoByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockIRepository(ctrl)
	repository.NewStore(m)

	repo := &model.RepoInfo{

		ID:          1,
		Name:        "houdini",
		CreatedAt:   "2021-01-01T00:00:00Z",
		UpdatedAt:   "2021-01-01T00:00:00Z",
		URL:         "test-url.com",
		Description: "test description",
		Language:    "Go",
		Forks:       10,
		Stars:       10,
		OpenIssues:  10,
	}
	m.EXPECT().GetRepoByName(gomock.Any(), "houdini").Return(repo, nil).AnyTimes()
	testCases := []struct {
		name       string
		pattern    string
		path       string
		assertions func(res *response.HttpResponse, code int, err error)
	}{
		{
			name:    "get repos by name success",
			pattern: "/v1/repo/{name}",
			path:    "/v1/repo/houdini",
			assertions: func(res *response.HttpResponse, code int, err error) {
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, code)
				assert.Equal(t, "repo retrieved successfully", res.Message)
				data := res.Data.(interface{})
				repo := data.(map[string]interface{})
				assert.Equal(t, float64(1), repo["id"])
				assert.Equal(t, "houdini", repo["name"])
				assert.Equal(t, "2021-01-01T00:00:00Z", repo["created_at"])
				assert.Equal(t, "2021-01-01T00:00:00Z", repo["updated_at"])
				assert.Equal(t, "test description", repo["description"])
				assert.Equal(t, "Go", repo["language"])
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, code, err := executeGetRequest(tc.pattern, tc.path, handler.GetRepoByName)
			tc.assertions(res, code, err)
		})

	}
}
