package main

import (
	"github.com/asavt7/nixEducation/pkg/configs"
	"github.com/asavt7/nixEducation/pkg/server"
	"github.com/asavt7/nixEducation/pkg/service"
	"github.com/asavt7/nixEducation/pkg/storage"
	"github.com/asavt7/nixEducation/pkg/tokenstorage"
	"time"
)

func main() {

	db := storage.NewPostgreDb(configs.InitPostgresConfig())
	store := storage.NewPostgresStorage(db)

	redisCacheStore := tokenstorage.InitRedisClient(configs.InitRedisConf())
	tokenStore := tokenstorage.NewTokenStorage(redisCacheStore, 10*time.Minute)

	srvc := service.NewService(store, tokenStore)
	handler := server.NewApiHandler(srvc)
	srvr := server.NewApiServer(handler)
	srvr.Run()
}
