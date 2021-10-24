package model

import "fmt"

type PostNotFoundErr struct {
	ID int
}

func (err PostNotFoundErr) Error() string {
	return fmt.Sprintf("post id=%d not found", err.ID)
}

type CommentNotFoundErr struct {
	ID int
}

func (err CommentNotFoundErr) Error() string {
	return fmt.Sprintf("comment id=%d not found", err.ID)
}

type UserNotFoundErr struct {
	ID int
}

func (err UserNotFoundErr) Error() string {
	return fmt.Sprintf("user id=%d not found", err.ID)
}

type UserHasNoAccessToChangeComment struct {
	UserID    int
	CommentID int
}

func (err UserHasNoAccessToChangeComment) Error() string {
	return fmt.Sprintf("user id=%d cannot change comment id=%d", err.UserID, err.CommentID)
}
