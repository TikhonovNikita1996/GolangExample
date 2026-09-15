package model

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrLoginFailed     = errors.New("login failed")
	ErrExternal        = errors.New("external error")
	ErrSessionNotFound = errors.New("session not found")
)
