package github

import (
	"context"
	"github.com/dilly3/houdini/internal/config"
	errs "github.com/dilly3/houdini/internal/error"
	"github.com/dilly3/houdini/internal/model"
	"github.com/dilly3/houdini/internal/repository"
	"github.com/dilly3/houdini/internal/repository/cache"
	"github.com/dilly3/houdini/pkg/github"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

func (g *GHubITR) ListCommits(owner, repo, since string, page int) ([]model.CommitInfo, error) {
	var commitsInfo []model.CommitInfo
	perPage := cache.GetDefaultCache().GetPerPage()
	perP, err := strconv.Atoi(perPage)
	if err != nil {
		return nil, errs.NewAppError("ListCommits:failed to convert perPage to int", err)
	}

	res, err := g.ghc.ListCommits(owner, repo, since, perP, page)
	if err != nil {
		return nil, errs.NewAppError("ListCommits:failed to decode commits,", err)
	}
	if len(res) < 1 {
		return commitsInfo, nil
	}
	commitsInfo = mapToCommitsInfo(res, repo)
	return commitsInfo, nil
}

// GetCommitsCron runs in the background to fetch commits
func (g *GHubITR) GetCommitsCron() error {
	var since *string
	cac := cache.GetDefaultCache()
	store := repository.GetDefaultStore()
	cmt, err := store.GetLastCommit(context.Background(), cac.GetRepo())
	if err != nil {
		s := cache.GetDefaultCache().GetSince()
		since = &s
	} else {
		since = &cmt.Date
	}
	perP, err := strconv.Atoi(cac.GetPerPage())

	if err != nil {
		return errs.NewAppError("GetCommitsCron:failed to convert perPage to int", err)
	}
	log.Info().Msg("fetching commits for repo:: " + cac.GetRepo())
	completeChan := make(chan bool)
	responseChan := make(chan []github.CommitResponse)
	tm := config.GetTimeDuration()
	// fetch commits in the background
	go func(chan bool, chan []github.CommitResponse, []int, *string, time.Duration) {
		startTime := time.Now()
		retries := config.Config.NetworkRetry
		for {
			log.Info().Msg("fetching commits page:: " + strconv.Itoa(g.page[0]))
			time.Sleep(10 * time.Second)
			res, err := g.ghc.ListCommits(cac.GetOwner(), cac.GetRepo(), *since, perP, g.page[0])
			if err != nil {
				log.Error().Err(err).Msg("failed to get commits")
				retries--
				log.Printf("retring: retry left %d", retries)
				if retries < 1 {
					completeChan <- true
					break
				}
				continue
			}
			responseChan <- res
			if len(res) < 1 {
				// reset page number to 1 because we have successfully fetched all commits
				g.mu.Lock()
				g.page[0] = 1
				g.mu.Unlock()
				completeChan <- true
				break
			}
			// terminate the goroutine before a new cron starts
			if time.Since(startTime) > (tm - time.Minute) {
				// increase page number by 1 before the end of the goroutine
				g.mu.Lock()
				g.page[0]++
				g.mu.Unlock()
				completeChan <- true
				break
			}
			// increase page number by 1
			g.mu.Lock()
			g.page[0]++
			g.mu.Unlock()

		}
		return

	}(completeChan, responseChan, g.page, since, tm)

	// listen for the response from fetched commits
	for {
		select {
		case <-completeChan:
			break
		case res := <-responseChan:
			var commitsSlice []model.CommitInfo
			if len(res) < 1 {
				break
			}
			commitsSlice = mapToCommitsInfo(res, cac.GetRepo())
			ctx := context.Background()
			err = repository.GetDefaultStore().SaveCommits(ctx, commitsSlice)
			if err != nil {
				log.Error().Err(err).Msg("failed to save commit")
				return errs.NewAppError("GetCommitsCron:failed to save commit", err)
			}
		}
	}

}
