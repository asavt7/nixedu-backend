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
	remoteUrl url.URL
	client    http.Client
}

func (r *RemoteService) GetAllByPostId(id int) ([]model.Comment, error) {
	nu := r.remoteUrl
	nu.Path = path.Join(r.remoteUrl.Path, "comments")

	req, err := http.NewRequest("GET", nu.String(), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("postId", strconv.Itoa(id))
	req.URL.RawQuery = q.Encode()

	response, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var comments []model.Comment

	err = json.NewDecoder(response.Body).Decode(&comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *RemoteService) GetAllByUserId(userId int) ([]model.Post, error) {
	nu := r.remoteUrl
	nu.Path = path.Join(r.remoteUrl.Path, "posts")

	req, err := http.NewRequest("GET", nu.String(), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("userId", strconv.Itoa(userId))
	req.URL.RawQuery = q.Encode()

	response, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var posts []model.Post

	err = json.NewDecoder(response.Body).Decode(&posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func NewRemoteService(remoteUrl string) *RemoteService {

	u, err := url.Parse(remoteUrl)
	if err != nil {
		log.Fatal(err)
	}
	return &RemoteService{
		remoteUrl: *u,
		client:    http.Client{},
	}
}

func (r *RemoteService) GetAll() ([]model.Post, error) {
	u := r.remoteUrl
	u.Path = path.Join(r.remoteUrl.Path, "posts")

	wg := &sync.WaitGroup{}

	resultCh := make(chan model.Post)
	wg.Add(100)
	for i := 1; i <= 100; i++ {
		go func(postId int, wg *sync.WaitGroup, resultCh chan model.Post) {
			nu := u
			nu.Path = path.Join(r.remoteUrl.Path, "posts", strconv.Itoa(postId))
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
