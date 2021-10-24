package main

import (
	"github.com/asavt7/nixedu/backend/internal/config"
	"github.com/asavt7/nixedu/backend/internal/server"
	"github.com/asavt7/nixedu/backend/internal/service"
	"github.com/asavt7/nixedu/backend/internal/storage"
	"github.com/asavt7/nixedu/backend/internal/tokenstorage"
	"time"
)

// @title NIX_EDU_APP
// @version 1.0
// @description This is a backend for nix tasks https://education.nixsolutions.com/mod/page/view.php?id=79

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config.InitConfigs()

	db := storage.NewPostgreDb(config.InitPostgresConfig())
	store := storage.NewPostgresStorage(db)

	redisCacheStore := tokenstorage.InitRedisClient(config.InitRedisConf())
	tokenStore := tokenstorage.NewTokenStorage(redisCacheStore, 10*time.Minute)

	srvc := service.NewService(store, tokenStore)
	handler := server.NewAPIHandler(srvc)
	srvr := server.NewAPIServer(handler)
	srvr.Run()
}
