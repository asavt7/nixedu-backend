package main

import (
	"github.com/asavt7/nixEducation/pkg/configs"
	"github.com/asavt7/nixEducation/pkg/server"
	"github.com/asavt7/nixEducation/pkg/service"
	"github.com/asavt7/nixEducation/pkg/storage"
	"github.com/asavt7/nixEducation/pkg/tokenstorage"
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

	db := storage.NewPostgreDb(configs.InitPostgresConfig())
	store := storage.NewPostgresStorage(db)

	redisCacheStore := tokenstorage.InitRedisClient(configs.InitRedisConf())
	tokenStore := tokenstorage.NewTokenStorage(redisCacheStore, 10*time.Minute)

	srvc := service.NewService(store, tokenStore)
	handler := server.NewApiHandler(srvc)
	srvr := server.NewApiServer(handler)
	srvr.Run()
}
