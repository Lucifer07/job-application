package util

import "errors"

var (
	ErrorBadRequest     = errors.New("error bad request")
	ErrorInternal       = errors.New("internal server error")
	ErrorUserNotFound   = errors.New("user not found")
	ErrorWrongPassword  = errors.New("wrong password")
	ErrorUnauthorized   = errors.New("unauthorized")
	ErrorInvalidToken   = errors.New("invalid token")
	ErrorForbidden      = errors.New("forbidden access")
	ErrorJobNotFound    = errors.New("job not found")
	ErrorJobfull        = errors.New("job full")
	ErrorAlreadyApplied = errors.New("already applied job")
)
