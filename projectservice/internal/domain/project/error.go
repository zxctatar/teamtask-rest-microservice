package projectdomain

import "errors"

var (
	ErrInvalidOwnerId = errors.New("invalid owner id")
	ErrInvalidName    = errors.New("invalid name")
)
