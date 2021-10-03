package service

import (
	"encoding/json"
	"github.com/asavt7/nixEducation/pkg/model"
	"log"
	"net/http"
	"net/url"
	"path"
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

	response, err := http.Get(u.String())
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
