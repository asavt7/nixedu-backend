package service

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/asavt7/nixEducation/pkg/storage"
	"time"
)

type PostService interface {
	GetAll() ([]model.Post, error)
	GetAllByUserId(userId int) ([]model.Post, error)
}

type CommentService interface {
	GetAllByPostId(id int) ([]model.Comment, error)
}

type UserService interface {
	CreateUser(user model.User, password string) (model.User, error)
	GetUser(id int) (model.User, error)
}

type AuthorizationService interface {
	CheckUserCredentials(username string, password string) (model.User, error)
	GenerateTokens(userId int) (accessToken, refreshToken string, accessExp, refreshExp time.Time, err error)
	ParseTokenToClaims(token string) (Claims, error)
	IsNeedToRefresh(claims Claims) bool
}

type Service struct {
	storage *storage.Storage

	AuthorizationService
	UserService
	PostService
	CommentService
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		storage:              storage,
		AuthorizationService: &AuthorizationServiceImpl{repo: storage.UserStorage},
		UserService:          nil,
		PostService:          nil,
		CommentService:       nil,
	}
}
