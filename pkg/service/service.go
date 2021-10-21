package service

import (
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/asavt7/nixEducation/pkg/storage"
	"github.com/asavt7/nixEducation/pkg/tokenstorage"
	"time"
)

type PostService interface {
	GetAll() ([]model.Post, error)
	GetAllByUserID(userID int) ([]model.Post, error)
	GetByUserIDAndPostID(userID, postID int) (model.Post, error)
	Save(userID int, post model.Post) (model.Post, error)
	Update(userID, postID int, updatePost model.UpdatePost) (model.Post, error)
	DeletePost(userID, postID int) error
}

type CommentService interface {
	GetAllByPostID(postID int) ([]model.Comment, error)
	Save(comment model.Comment) (model.Comment, error)
	Update(currentUserID, commentID int, comment model.UpdateComment) (model.Comment, error)
	Delete(currentUserID, commentID int) error
}

type UserService interface {
	CreateUser(user model.User, password string) (model.User, error)
	GetUserById(id int) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
}

// AuthorizationService interface contains methods for working with tokens
type AuthorizationService interface {
	CheckUserCredentials(username string, password string) (model.User, error)
	GenerateTokens(userID int) (accessToken, refreshToken string, accessExp, refreshExp time.Time, err error)
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
		UserService: &UserServiceImpl{
			repo: storage.UserStorage,
		},
		PostService:    nil,
		CommentService: nil,
	}
}
