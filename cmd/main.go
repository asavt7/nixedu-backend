package main

import (
	"github.com/asavt7/nixEducation/pkg/server"
	"github.com/asavt7/nixEducation/pkg/service"
	"github.com/asavt7/nixEducation/pkg/storage"
	"log"
)

func main() {

	db, err := storage.NewPostgreDb(storage.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal(err)
	}

	store := storage.NewPostgresStorage(db)

	srvc := service.NewService(store)
	handler := server.NewApiHandler(srvc)
	srvr := server.NewApiServer(handler)
	srvr.Run()
}
