package models

import "errors"

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound = errors.New("not found")
	ErrInternal = errors.New("internal repository error")
)
