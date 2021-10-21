package storage_test

import (
	"errors"
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/asavt7/nixEducation/pkg/storage"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func initUserStorage(t *testing.T) (storage.UserStorage, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := storage.NewPostgresStorage(db)
	return repo, mock, func() {
		_ = db.Close()
	}
}

func TestPostgresUserStorage_Create(t *testing.T) {
	repo, mock, destroyFunc := initUserStorage(t)
	defer destroyFunc()

	tests := []struct {
		name    string
		s       storage.UserStorage
		user    model.User
		mock    func()
		want    model.User
		wantErr bool
	}{
		{
			name: "OK",
			s:    repo,
			user: model.User{
				Email:        "email@googlle.com",
				Username:     "username",
				PasswordHash: "password",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO nix.users \\(username,email,password_hash\\) values").WithArgs("username", "email@googlle.com", "password").WillReturnRows(rows)
			},
			want: model.User{
				ID:           1,
				Username:     "username",
				Email:        "email@googlle.com",
				PasswordHash: "password",
			},
		},
		{
			name: "empty fields",
			s:    repo,
			user: model.User{},
			mock: func() {
				mock.ExpectQuery("INSERT INTO nix.users \\(username,email,password_hash\\) values").WithArgs("", "", "").WillReturnError(errors.New("invalid values"))
			},
			want:    model.User{},
			wantErr: true,
		},
		{
			name: "user already exists",
			s:    repo,
			user: model.User{
				Email:        "email@googlle.com",
				Username:     "username",
				PasswordHash: "password",
			},
			mock: func() {
				mock.ExpectQuery("INSERT INTO nix.users \\(username,email,password_hash\\) values").WithArgs("username", "email@googlle.com", "password").WillReturnError(errors.New("user already exists"))
			},
			want:    model.User{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Create(tt.user)
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

func TestPostgresUserStorage_FindByEmail(t *testing.T) {
	repo, mock, destroyFunc := initUserStorage(t)
	defer destroyFunc()

	tests := []struct {
		name    string
		s       storage.UserStorage
		user    model.User
		mock    func()
		wantErr bool
	}{
		{
			name: "OK",
			s:    repo,
			user: model.User{
				ID:           1,
				Email:        "email@googlle.com",
				Username:     "username",
				PasswordHash: "password",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "email", "password_hash"}).AddRow(1, "username", "email@googlle.com", "password")
				mock.ExpectQuery("SELECT \\* FROM nix.users WHERE email").WithArgs("email@googlle.com").WillReturnRows(rows)
			},
		},
		{
			name: "no user found",
			s:    repo,
			user: model.User{
				ID:           1,
				Email:        "email@googlle.com",
				Username:     "username",
				PasswordHash: "password",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "email", "password_hash"})
				mock.ExpectQuery("SELECT \\* FROM nix.users WHERE email").WithArgs("email@googlle.com").WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.FindByEmail(tt.user.Email)
			if ((err != nil) != tt.wantErr) && (err != model.UserNotFoundErr{ID: 0}) {
				t.Errorf(" error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.user {
				t.Errorf(" = %v, want %v", got, tt.user)
			}
		})
	}

}

func TestPostgresUserStorage_GetById(t *testing.T) {
	repo, mock, destroyFunc := initUserStorage(t)
	defer destroyFunc()

	tests := []struct {
		name    string
		s       storage.UserStorage
		user    model.User
		mock    func()
		wantErr bool
	}{
		{
			name: "OK",
			s:    repo,
			user: model.User{
				ID:           1,
				Email:        "email@googlle.com",
				Username:     "username",
				PasswordHash: "password",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "email", "password_hash"}).AddRow(1, "username", "email@googlle.com", "password")
				mock.ExpectQuery("SELECT \\* FROM nix.users WHERE id").WithArgs(1).WillReturnRows(rows)
			},
		},
		{
			name: "no user found",
			s:    repo,
			user: model.User{
				ID:           1,
				Email:        "email@googlle.com",
				Username:     "username",
				PasswordHash: "password",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "email", "password_hash"})
				mock.ExpectQuery("SELECT \\* FROM nix.users WHERE id").WithArgs(1).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.GetByID(tt.user.ID)
			if ((err != nil) != tt.wantErr) && (err != model.UserNotFoundErr{ID: tt.user.ID}) {
				t.Errorf(" error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.user {
				t.Errorf(" = %v, want %v", got, tt.user)
			}
		})
	}
}

func TestPostgresUserStorage_FindByUsername(t *testing.T) {
	repo, mock, destroyFunc := initUserStorage(t)
	defer destroyFunc()

	tests := []struct {
		name    string
		s       storage.UserStorage
		user    model.User
		mock    func()
		wantErr bool
	}{
		{
			name: "OK",
			s:    repo,
			user: model.User{
				ID:           1,
				Email:        "email@googlle.com",
				Username:     "username",
				PasswordHash: "password",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "email", "password_hash"}).AddRow(1, "username", "email@googlle.com", "password")
				mock.ExpectQuery("SELECT \\* FROM nix.users WHERE username").WithArgs("username").WillReturnRows(rows)
			},
		},
		{
			name: "no user found",
			s:    repo,
			user: model.User{
				ID:           1,
				Email:        "email@googlle.com",
				Username:     "username",
				PasswordHash: "password",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "email", "password_hash"})
				mock.ExpectQuery("SELECT \\* FROM nix.users WHERE username").WithArgs("username").WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.FindByUsername(tt.user.Username)
			if ((err != nil) != tt.wantErr) && (err != model.UserNotFoundErr{ID: 0}) {
				t.Errorf(" error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.user {
				t.Errorf(" = %v, want %v", got, tt.user)
			}
		})
	}
}

func TestPostgresUserStorage_GetByUsernameAndPasswordHash(t *testing.T) {
	repo, mock, destroyFunc := initUserStorage(t)
	defer destroyFunc()

	tests := []struct {
		name    string
		s       storage.UserStorage
		user    model.User
		mock    func()
		wantErr bool
	}{
		{
			name: "OK",
			s:    repo,
			user: model.User{
				ID:           1,
				Email:        "email@googlle.com",
				Username:     "username",
				PasswordHash: "password",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "email", "password_hash"}).AddRow(1, "username", "email@googlle.com", "password")
				mock.ExpectQuery("SELECT \\* FROM nix.users WHERE username=\\$1 and password_hash=\\$2").WithArgs("username", "password").WillReturnRows(rows)
			},
		},
		{
			name: "no user found",
			s:    repo,
			user: model.User{
				ID:           1,
				Email:        "email@googlle.com",
				Username:     "username",
				PasswordHash: "password",
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "username", "email", "password_hash"})
				mock.ExpectQuery("SELECT \\* FROM nix.users WHERE username=\\$1 and password_hash=\\$2").WithArgs("username", "password").WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.GetByUsernameAndPasswordHash(tt.user.Username, tt.user.PasswordHash)
			if ((err != nil) != tt.wantErr) && (err != model.UserNotFoundErr{ID: 0}) {
				t.Errorf(" error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got != tt.user {
				t.Errorf(" = %v, want %v", got, tt.user)
			}
		})
	}
}
