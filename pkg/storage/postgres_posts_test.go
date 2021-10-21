package storage_test

import (
	"errors"
	"github.com/asavt7/nixedu/backend/pkg/model"
	"github.com/asavt7/nixedu/backend/pkg/storage"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func initPostsStorage(t *testing.T) (storage.PostsStorage, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := storage.NewPostgresStorage(db).PostsStorage
	return repo, mock, func() {
		_ = db.Close()
	}
}

func TestPostgresPostsStorage_Save(t *testing.T) {
	repo, mock, destroyFunc := initPostsStorage(t)
	defer destroyFunc()

	tests := []struct {
		name    string
		s       storage.PostsStorage
		post    model.Post
		mock    func()
		want    model.Post
		wantErr bool
	}{
		{
			name: "OK",
			s:    repo,
			post: model.Post{
				UserID: 1,
				Title:  "title",
				Body:   "body",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO nix.posts \\(userid, title, body\\) VALUES").WithArgs(1, "title", "body").WillReturnRows(rows)
			},
			want: model.Post{
				UserID: 1,
				ID:     1,
				Title:  "title",
				Body:   "body",
			},
		},
		{
			name: "empty fields",
			s:    repo,
			post: model.Post{},
			mock: func() {
				mock.ExpectQuery("INSERT INTO nix.posts \\(userid, title, body\\) VALUES").WithArgs(0, "", "").WillReturnError(errors.New("invalid values"))
			},
			want:    model.Post{},
			wantErr: true,
		},
		{
			name: "user already exists",
			s:    repo,
			post: model.Post{
				UserID: 1,
				Title:  "title",
				Body:   "body",
			},
			mock: func() {
				mock.ExpectQuery("INSERT INTO nix.posts \\(userid, title, body\\) VALUES").WithArgs(1, "title", "body").WillReturnError(errors.New("user already exists"))
			},
			want: model.Post{
				UserID: 1,
				Title:  "title",
				Body:   "body",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Save(tt.post)
			if (err != nil) != tt.wantErr {
				t.Errorf(" error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf(" = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresPostsStorage_GetAll(t *testing.T) {
	repo, mock, destroyFunc := initPostsStorage(t)
	defer destroyFunc()

	tests := []struct {
		name    string
		s       storage.PostsStorage
		mock    func()
		want    []model.Post
		wantErr bool
	}{
		{
			name: "OK",
			s:    repo,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "userid", "title", "body"}).AddRow(1, 1, "title", "body")
				mock.ExpectQuery("SELECT \\* FROM nix.posts").WillReturnRows(rows)
			},
			want: []model.Post{{
				UserID: 1,
				ID:     1,
				Title:  "title",
				Body:   "body"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf(" error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil || !assert.EqualValues(t, tt.want, got) {
				t.Errorf(" = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresPostsStorage_GetAllByUserID(t *testing.T) {
	repo, mock, destroyFunc := initPostsStorage(t)
	defer destroyFunc()

	tests := []struct {
		name    string
		s       storage.PostsStorage
		mock    func()
		want    []model.Post
		wantErr bool
	}{
		{
			name: "OK",
			s:    repo,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "userid", "title", "body"}).AddRow(1, 1, "title", "body")
				mock.ExpectQuery("SELECT \\* FROM nix.posts WHERE userid=\\$1").WithArgs(1).WillReturnRows(rows)
			},
			want: []model.Post{{
				UserID: 1,
				ID:     1,
				Title:  "title",
				Body:   "body"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.GetAllByUserID(1)
			if (err != nil) != tt.wantErr {
				t.Errorf(" error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !assert.Equal(t, tt.want, got) {
				t.Errorf(" = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresPostsStorage_Update(t *testing.T) {
	repo, mock, destroyFunc := initPostsStorage(t)
	defer destroyFunc()

	title := "title"
	body := "body"

	tests := []struct {
		name    string
		s       storage.PostsStorage
		post    model.UpdatePost
		mock    func()
		want    model.Post
		wantErr bool
	}{
		{
			name: "OK",
			s:    repo,
			post: model.UpdatePost{
				Title: &title,
				Body:  &body,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "userid", "title", "body"}).AddRow(1, 1, "title", "body")
				mock.ExpectQuery("UPDATE nix.posts SET title=\\$1, body=\\$2 WHERE userid=\\$3 AND id=\\$4").WithArgs(title, body, 1, 1).WillReturnRows(rows)
			},
			want: model.Post{
				UserID: 1,
				ID:     1,
				Title:  "title",
				Body:   "body",
			},
		},
		{
			name: "OK only title",
			s:    repo,
			post: model.UpdatePost{
				Title: &title,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "userid", "title", "body"}).AddRow(1, 1, "title", "body")
				mock.ExpectQuery("UPDATE nix.posts SET title=\\$1 WHERE userid=\\$2 AND id=\\$3").WithArgs(title, 1, 1).WillReturnRows(rows)
			},
			want: model.Post{
				UserID: 1,
				ID:     1,
				Title:  "title",
				Body:   "body",
			},
		},
		{
			name: "OK only body",
			s:    repo,
			post: model.UpdatePost{
				Body: &body,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "userid", "title", "body"}).AddRow(1, 1, "title", "body")
				mock.ExpectQuery("UPDATE nix.posts SET body=\\$1 WHERE userid=\\$2 AND id=\\$3").WithArgs(body, 1, 1).WillReturnRows(rows)
			},
			want: model.Post{
				UserID: 1,
				ID:     1,
				Title:  "title",
				Body:   "body",
			},
		},
		{
			name: "error",
			s:    repo,
			post: model.UpdatePost{
				Body: &body,
			},
			mock: func() {
				mock.ExpectQuery("UPDATE nix.posts SET body=\\$1 WHERE userid=\\$2 AND id=\\$3").WithArgs(body, 1, 1).WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "error - no update params provided",
			s:    repo,
			post: model.UpdatePost{},
			mock: func() {
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Update(1, 1, tt.post)
			if (err != nil) != tt.wantErr {
				t.Errorf(" error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.want {
				t.Errorf(" = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostgresPostsStorage_DeleteByUserIDAndID(t *testing.T) {
	repo, mock, destroyFunc := initPostsStorage(t)
	defer destroyFunc()

	tests := []struct {
		name    string
		s       storage.PostsStorage
		post    model.Post
		mock    func()
		want    model.Post
		wantErr bool
	}{
		{
			name: "OK",
			s:    repo,
			post: model.Post{
				UserID: 1, ID: 1,
				Title: "title",
				Body:  "body",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("DELETE FROM nix.posts").WithArgs(1, 1).WillReturnRows(rows)
			},
		},
		{
			name: "not found",
			s:    repo,
			post: model.Post{},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("DELETE FROM nix.posts").WithArgs(1, 1).WillReturnRows(rows)
			},
			want:    model.Post{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := tt.s.DeleteByUserIDAndID(1, 1)
			if (err != nil) != tt.wantErr {
				t.Errorf(" error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
