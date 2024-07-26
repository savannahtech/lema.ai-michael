package postgres

import (
	"fmt"
	"github.com/dilly3/houdini/internal/config"
	model2 "github.com/dilly3/houdini/internal/model"
	"github.com/dilly3/houdini/internal/repository"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Storage object
type Storage struct {
	DB     *gorm.DB
	Logger *zerolog.Logger
}

var DefaultStorage *Storage

func New(config *config.Configuration, logger *zerolog.Logger) *repository.Store {

	postgresDSN := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=%s",
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUser,
		config.PostgresDB,
		config.PostgresPassword,
		config.PostgresTimezone,
	)
	db, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{})
	if err != nil {
		logger.Error().Err(err).Msg("failed to connect to db")

	}
	err = db.AutoMigrate(&model2.CommitInfo{}, &model2.RepoInfo{})
	if err != nil {
		panic(err)
	}
	logger.Info().Msg("connected to db")
	str := &Storage{
		DB:     db,
		Logger: logger,
	}
	if DefaultStorage == nil {
		DefaultStorage = str
	}
	cs := NewCommitStore(str)
	rs := NewRepoStore(str)
	pgs := &repository.Store{
		ICommitRepository: cs, IRepoRepository: rs,
	}

	return pgs

}
