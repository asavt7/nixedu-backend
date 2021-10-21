package storage

import (
	"encoding/json"
	"github.com/asavt7/nixEducation/pkg/model"
	"io/ioutil"
	"path"
	"strconv"
)

type FsStorage struct {
	basePath string
}

func NewFsStorage(basePath string) *FsStorage {
	return &FsStorage{basePath: basePath}
}

func (f *FsStorage) SaveAll(posts []model.Post) ([]model.Post, error) {

	for _, post := range posts {
		pathToFile := path.Join(f.basePath, "posts", strconv.Itoa(post.ID)+".txt")

		content, err := json.MarshalIndent(post, "", "    ")
		if err != nil {
			return nil, err
		}

		err = ioutil.WriteFile(pathToFile, []byte(content), 0644)
		if err != nil {
			return nil, err
		}
	}
	return posts, nil
}
