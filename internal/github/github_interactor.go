package github

import (
	"github.com/dilly3/houdini/pkg/github"
	"sync"
)

var gitHubItr *GHubITR

type GHubITR struct {
	ghc  *github.GHClient
	page []int
	mu   sync.Mutex
}

func setGitHubAdp(adaptor *GHubITR) {
	gitHubItr = adaptor
}

func GetGitHubAdp() *GHubITR {
	return gitHubItr
}

// NewGHubITR sets up new github interactor
func NewGHubITR(ghc *github.GHClient) *GHubITR {
	ghi := &GHubITR{
		ghc,
		[]int{1},
		sync.Mutex{},
	}
	setGitHubAdp(ghi)
	return ghi
}
