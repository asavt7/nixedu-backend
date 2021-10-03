package main

import (
	"github.com/asavt7/nixEducation/pkg/service"
	"github.com/asavt7/nixEducation/pkg/storage"
	"log"
)

func main() {

	srv := service.NewRemoteService("https://jsonplaceholder.typicode.com/")
	stora := storage.NewFsStorage("./storage/")

	posts, err := srv.GetAll()

	if err != nil {
		log.Fatal(err)
	}
	_, err = stora.SaveAll(posts)
	if err != nil {
		log.Fatal(err)
	}

}
