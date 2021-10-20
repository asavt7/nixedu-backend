package model

import "fmt"

type PostNotFoundErr struct {
	Id int
}

func (err PostNotFoundErr) Error() string {
	return fmt.Sprintf("post id=%d not found", err.Id)
}

type CommentNotFoundErr struct {
	Id int
}

func (err CommentNotFoundErr) Error() string {
	return fmt.Sprintf("comment id=%d not found", err.Id)
}

type UserNotFoundErr struct {
	Id int
}

func (err UserNotFoundErr) Error() string {
	return fmt.Sprintf("user id=%d not found", err.Id)
}

type UserHasNoAccessToChangeComment struct {
	UserId    int
	CommentId int
}

func (err UserHasNoAccessToChangeComment) Error() string {
	return fmt.Sprintf("user id=%d cannot change comment id=%d", err.UserId, err.CommentId)
}
