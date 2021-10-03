package main

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/asavt7/nixEducation/pkg/service"
	"github.com/asavt7/nixEducation/pkg/storage"
	"log"
	"sync"
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

	srv := service.NewRemoteService("https://jsonplaceholder.typicode.com/")
	store := storage.NewPostgresStorage(db)

	posts, err := srv.GetAllByUserId(7)
	if err != nil {
		log.Fatal(err)
	}

	wgp := &sync.WaitGroup{}
	for _, post := range posts {
		wgp.Add(1)
		go func(post model.Post, wgp *sync.WaitGroup) {
			post, err := store.PostsStorage.Save(post)
			if err != nil {
				log.Fatalf("error saving post %s", err.Error())
			}

			comments, err := srv.GetAllByPostId(post.Id)
			if err != nil {
				log.Fatal(err)
			}

			wgc := &sync.WaitGroup{}

			for _, comment := range comments {
				wgc.Add(1)
				go func(comment model.Comment, wgc *sync.WaitGroup) {
					comment, err = store.CommentsStorage.Save(comment)
					if err != nil {
						log.Fatalf("error saving comment %s", err.Error())
					}
					wgc.Done()
				}(comment, wgc)
			}
			wgc.Wait()
			wgp.Done()
		}(post, wgp)
	}
	wgp.Wait()

}
