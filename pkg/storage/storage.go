package storage

import "github.com/asavt7/nixEducation/pkg/model"

type PostsStorage interface {
	Save(p model.Post) (model.Post, error)
	GetAll() ([]model.Post, error)
	GetAllByUserID(userID int) ([]model.Post, error)
	GetByUserIDAndID(userID, postID int) (model.Post, error)

	Update(userID, postID int, p model.UpdatePost) (model.Post, error)
	DeleteByUserIDAndID(userID, postID int) error
}

type CommentsStorage interface {
	SaveAll(posts []model.Comment) ([]model.Comment, error)
	Save(c model.Comment) (model.Comment, error)

	GetAllByUserID(userID int) ([]model.Comment, error)
	GetAllByPostID(postID int) ([]model.Comment, error)
	GetByCommentID(commentID int) (model.Comment, error)

	Update(userID, commentID int, p model.UpdateComment) (model.Comment, error)
	DeleteByUserIDAndID(userID, commentID int) error
}

type UserStorage interface {
	GetByUsernameAndPasswordHash(username, passwordHash string) (model.User, error)
	Create(user model.User) (model.User, error)
	GetByID(userID int) (model.User, error)
	FindByUsername(username string) (model.User, error)
	FindByEmail(email string) (model.User, error)
}

type Storage struct {
	PostsStorage
	CommentsStorage
	UserStorage
}
