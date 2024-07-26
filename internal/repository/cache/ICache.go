package cache

var defaultCache ICache

type ICache interface {
	GetOwner() string
	SetOwner(owner string)
	GetRepo() string
	SetRepo(repo string)
	GetSince() string
	SetSince(since string)
	GetPerPage() string
	SetPerPage(perPage string)
	Close() error
}

func setCache(cache ICache) {
	defaultCache = cache
}
func GetDefaultCache() ICache {
	return defaultCache
}
func NewCache(cac ICache) ICache {
	setCache(cac)
	return cac
}
