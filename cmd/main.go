package main

import (
	"encoding/json"
	"github.com/asavt7/nixEducation/pkg/service"
	"log"
	"os"
)

func main() {

	srv := service.NewRemoteService("https://jsonplaceholder.typicode.com/")

	posts, err := srv.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	e := json.NewEncoder(os.Stdout)
	e.SetIndent("    ", "    ")

	err = e.Encode(posts)
	if err != nil {
		log.Fatal(err)
	}

}
