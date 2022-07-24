package errservice

import "errors"

var (
	ErrForbidden = errors.New("you don't have permission to access this resouce")
)