package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

var (
	owner   = "github_owner"
	repo    = "github_repo"
	since   = "github_since"
	perPage = "github_per_page"
)

type RdsClient struct {
	rd            *redis.Client
	logger        *zerolog.Logger
	defaultValues map[string]string
}

func NewRedisClient(addr, user, password string, deft map[string]string, logger *zerolog.Logger) (*RdsClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: user,
		Password: password,
		DB:       0,
	})

	s, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	logger.Info().Msgf("Redis connection: %s", s)

	return &RdsClient{
		rd:            rdb,
		logger:        logger,
		defaultValues: deft,
	}, nil
}

func (r *RdsClient) GetOwner() string {
	s, err := r.rd.Get(context.Background(), owner).Result()
	if err != nil {
		return r.defaultValues[owner]
	}
	return s
}

func (r *RdsClient) SetOwner(ownerName string) {
	r.rd.Set(context.Background(), owner, ownerName, 0)
}

func (r *RdsClient) GetRepo() string {
	s, err := r.rd.Get(context.Background(), repo).Result()
	if err != nil {
		return r.defaultValues[repo]
	}
	return s
}

func (r *RdsClient) SetRepo(repoName string) {
	r.rd.Set(context.Background(), repo, repoName, 0)
}

func (r *RdsClient) GetSince() string {
	s, err := r.rd.Get(context.Background(), since).Result()
	if err != nil {
		return r.defaultValues[since]
	}
	r.logger.Info().Msgf("Since: %s", s)
	return s
}

func (r *RdsClient) SetSince(sinceDate string) {
	if len(sinceDate) < 12 {
		sinceDateZ := sinceDate + "T00:00:00Z"
		r.rd.Set(context.Background(), since, sinceDateZ, 0)
	} else {
		r.rd.Set(context.Background(), since, sinceDate, 0)
	}

}

func (r *RdsClient) GetPerPage() string {
	s, err := r.rd.Get(context.Background(), perPage).Result()
	if err != nil {
		return r.defaultValues[perPage]
	}
	return s
}

func (r *RdsClient) SetPerPage(perP string) {
	r.rd.Set(context.Background(), perPage, perP, 0)
}

func (r *RdsClient) Close() error {
	return r.rd.Close()
}
