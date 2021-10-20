package storage

import "github.com/asavt7/nixEducation/pkg/model"

type PostsStorage interface {
	SaveAll(userId int, posts []model.Post) ([]model.Post, error)
	Save(userId int, p model.Post) (model.Post, error)

	GetAllByUserId(userId int) ([]model.Post, error)
	GetByUserIdAndId(userId, postId int) (model.Post, error)

	Update(userId, postId int, p model.UpdatePost) (model.Post, error)
	DeleteByUserIdAndId(userId, postId int) error
}

type CommentsStorage interface {
	SaveAll(posts []model.Comment) ([]model.Comment, error)
	Save(c model.Comment) (model.Comment, error)

	GetAllByUserId(userId int) ([]model.Comment, error)
	GetAllByPostId(postId int) ([]model.Comment, error)
	GetByCommentId(commentId int) (model.Comment, error)

	Update(userId, postId, commentId int, p model.UpdateComment) (model.Comment, error)
	DeleteByUserIdAndId(userId, commentId int) error
}

type UserStorage interface {
	GetByUsernameAndPasswordHash(username, passwordHash string) (model.User, error)
	Create(user model.User) (model.User, error)
	GetById(userId int) (model.User, error)
	FindByUsername(username string) (model.User, error)
	FindByEmail(email string) (model.User, error)
}

type Storage struct {
	PostsStorage
	CommentsStorage
	UserStorage
}
