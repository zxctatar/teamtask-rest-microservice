package deleteerr

import "errors"

var (
	ErrProjectNotFound  = errors.New("project not found")
	ErrInvalidProjectId = errors.New("invalid project id")
)
