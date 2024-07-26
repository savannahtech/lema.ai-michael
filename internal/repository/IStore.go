package repository

var defaultStore IRepository

func GetDefaultStore() IRepository {
	return defaultStore
}
func SetDefaultStore(store IRepository) {
	defaultStore = store
}

type IRepository interface {
	ICommitRepository
	IRepoRepository
}

type Store struct {
	ICommitRepository
	IRepoRepository
}

func NewStore(storage IRepository) IRepository {
	SetDefaultStore(storage)
	return storage
}
