package createerr

import "errors"

var (
	ErrAlreadyExists = errors.New("project already exists")
)
