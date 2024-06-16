package errors

import "errors"

var (
	ErrUnexpectedConnecion = errors.New("unexpected connection")
	ErrInvalidAuthToken    = errors.New("invalid auth token")
)
