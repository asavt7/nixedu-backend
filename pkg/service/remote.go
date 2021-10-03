package service

import (
	"encoding/json"
	"github.com/asavt7/nixEducation/pkg/model"
	"log"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strconv"
	"sync"
)

type RemoteService struct {
	RemoteUrl url.URL
}

func NewRemoteService(remoteUrl string) *RemoteService {

	u, err := url.Parse(remoteUrl)
	if err != nil {
		log.Fatal(err)
	}
	return &RemoteService{RemoteUrl: *u}
}

func (r *RemoteService) GetAll() ([]model.Post, error) {
	u := r.RemoteUrl
	u.Path = path.Join(r.RemoteUrl.Path, "posts")

	wg := &sync.WaitGroup{}

	resultCh := make(chan model.Post)
	wg.Add(100)
	for i := 1; i <= 100; i++ {
		go func(postId int, wg *sync.WaitGroup, resultCh chan model.Post) {
			nu := u
			nu.Path = path.Join(r.RemoteUrl.Path, "posts", strconv.Itoa(postId))
			response, err := http.Get(nu.String())
			if err != nil {
				log.Fatal(err)
			}

			var post model.Post

			err = json.NewDecoder(response.Body).Decode(&post)
			if err != nil {
				log.Fatal(err)
			}
			resultCh <- post
			wg.Done()
		}(i, wg, resultCh)
	}

	go func(wg *sync.WaitGroup, res chan model.Post) {
		wg.Wait()
		close(resultCh)
	}(wg, resultCh)

	posts := make([]model.Post, 0, 100)
	for r := range resultCh {
		posts = append(posts, r)
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Id < posts[j].Id
	})

	return posts, nil

}
