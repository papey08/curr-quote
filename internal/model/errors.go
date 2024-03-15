package model

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input body or path params")

	ErrInvalidCurr   = errors.New("given currency is invalid")
	ErrInvalidId     = errors.New("given id is invalid")
	ErrQuoteNotFound = errors.New("quote with required id does not exist")

	ErrDatabaseError = errors.New("something wrong with database")
	ErrApiError      = errors.New("something wrong with exchange api")
	ErrServiceError  = errors.New("something wrong with service")
)
