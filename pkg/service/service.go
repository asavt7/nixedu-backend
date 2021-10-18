package service

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/asavt7/nixEducation/pkg/storage"
	"github.com/asavt7/nixEducation/pkg/tokenstorage"
	"time"
)

type PostService interface {
	GetAll() ([]model.Post, error)
	GetAllByUserId(userId int) ([]model.Post, error)
	GetByUserIdAndPostId(userId, postId int) (model.Post, error)
	Save(userId int, post model.Post) (model.Post, error)
	Update(userId, postId int, updatePost model.UpdatePost) (model.Post, error)
	DeletePost(userId, postId int) error
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
	ParseAccessTokenToClaims(token string) (*Claims, error)
	ParseRefreshTokenToClaims(token string) (*Claims, error)
	IsNeedToRefresh(claims *Claims) bool
	Logout(accessTokenClaims *Claims) error
	ValidateAccessToken(accessTokenClaims *Claims) error
	ValidateRefreshToken(accessTokenClaims *Claims) error
}

type Service struct {
	storage *storage.Storage

	AuthorizationService
	UserService
	PostService
	CommentService
}

func NewService(storage *storage.Storage, tokenStore *tokenstorage.TokenStorage) *Service {
	return &Service{
		storage: storage,
		AuthorizationService: &AuthorizationServiceImpl{
			repo:       storage.UserStorage,
			tokenStore: tokenStore,
		},
		UserService:    nil,
		PostService:    nil,
		CommentService: nil,
	}
}
