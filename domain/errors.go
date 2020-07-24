package domain

import "errors"

var (
	// ErrUserWithEmailAlreadyExist error for duplicated email
	ErrUserWithEmailAlreadyExist = errors.New("user with email already exists")
	// ErrUserWithUsernameAlreadyExist error for duplicated usename
	ErrUserWithUsernameAlreadyExist = errors.New("user with usename already exists")
	// ErrNoResult not found
	ErrNoResult = errors.New("user was not found")
)
