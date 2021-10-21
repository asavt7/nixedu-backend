package tokenstorage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"time"
)

func InitRedisClient(opt *redis.Options) *redis.Client {
	client := redis.NewClient(opt)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}

	return client
}

type RedisTokenStore struct {
	client            *redis.Client
	AutoLogoffMinutes time.Duration
}

func NewRedisTokenStore(client *redis.Client, autoLogoffMinutes time.Duration) *RedisTokenStore {
	return &RedisTokenStore{client: client, AutoLogoffMinutes: autoLogoffMinutes}
}

func (r *RedisTokenStore) Get(userID int) (model.CachedTokens, error) {

	key := fmt.Sprintf("token-%s", strconv.Itoa(userID))
	res, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return model.CachedTokens{}, err
	}

	err = r.client.Expire(context.Background(), key, r.AutoLogoffMinutes).Err()
	if err != nil {
		return model.CachedTokens{}, err
	}

	cachedTokens := new(model.CachedTokens)
	err = json.Unmarshal([]byte(res), cachedTokens)
	return *cachedTokens, err
}

func (r *RedisTokenStore) Delete(userID int) error {
	key := fmt.Sprintf("token-%s", strconv.Itoa(userID))
	return r.client.Del(context.Background(), key).Err()
}

func (r *RedisTokenStore) Save(userID int, tokens model.CachedTokens) (model.CachedTokens, error) {
	key := fmt.Sprintf("token-%s", strconv.Itoa(userID))
	cacheJSON, err := json.Marshal(tokens)
	if err != nil {
		return tokens, err
	}
	err = r.client.Set(context.Background(), key, cacheJSON, r.AutoLogoffMinutes).Err()
	return tokens, err
}
