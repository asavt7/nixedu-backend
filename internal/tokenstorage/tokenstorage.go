package tokenstorage

import (
	"github.com/asavt7/nixedu/backend/internal/model"
	"github.com/go-redis/redis/v8"
	"time"
)

// TokenStorage struct
type TokenStorage struct {
	TokenKeeper
}

// NewTokenStorage construct new TokenStorage instance
func NewTokenStorage(client *redis.Client, autoLogoffMinutes time.Duration) *TokenStorage {
	return &TokenStorage{TokenKeeper: NewRedisTokenStore(client, autoLogoffMinutes)}
}

// TokenKeeper interface provides methods for storing model.CachedTokens in cache
type TokenKeeper interface {
	Get(userID int) (model.CachedTokens, error)
	Delete(userID int) error
	Save(userID int, tokens model.CachedTokens) (model.CachedTokens, error)
}
