package domain

import "errors"

var (
	ErrNotFound           = errors.New("data not found")
	FailedRegisterMessage = "failed to register user"
	ErrDataAlreadyExists = errors.New("data already exists")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrBadRequest        = errors.New("bad request")
	ErrTotalPageNotValid = errors.New("total page not valid")
	ErrReleaseYearNotValid = errors.New("release year not valid, must be between 1980 and 2024")
)
