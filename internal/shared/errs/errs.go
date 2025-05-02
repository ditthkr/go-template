package errs

import "errors"

var (
	ErrDuplicateUsername = errors.New("username already exists")
	ErrDuplicateEmail    = errors.New("email already exists")
	ErrInvalidCredential = errors.New("invalid credential")
)
