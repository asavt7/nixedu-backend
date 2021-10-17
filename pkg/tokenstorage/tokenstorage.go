package tokenstorage

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/go-redis/redis/v8"
	"time"
)

type TokenStorage struct {
	TokenKeeper
}

func NewTokenStorage(client *redis.Client, autoLogoffMinutes time.Duration) *TokenStorage {
	return &TokenStorage{TokenKeeper: NewRedisTokenStore(client, autoLogoffMinutes)}
}

type TokenKeeper interface {
	Get(userId int) (model.CachedTokens, error)
	Delete(userId int) error
	Save(userId int, tokens model.CachedTokens) (model.CachedTokens, error)
}
