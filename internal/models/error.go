package models

import "errors"

type Error struct {
	CustomError error
}

var (
	ErrNotFound     = Error{CustomError: errors.New("record not found")}
	ErrInternal     = Error{CustomError: errors.New("internal server error")}
	ErrUserNotFound = Error{CustomError: errors.New("user not found")}
)
