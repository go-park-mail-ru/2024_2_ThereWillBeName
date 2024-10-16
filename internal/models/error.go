package models

import "errors"

type Error struct {
	CustomError error
}

var (
	ErrUserAlreadyExists = Error{CustomError: errors.New("user already exists")}
)
