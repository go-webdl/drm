package remote

import "errors"

var (
	ErrNoServerURL     = errors.New("server URL not provided")
	ErrServerStatus    = errors.New("server status error")
	ErrInvalidResponse = errors.New("invalid response error")
)
