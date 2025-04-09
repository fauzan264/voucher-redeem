package constants

import "errors"

var (
	ErrWrongUserOrPassword = errors.New("Invalid email or password")
	ErrRecordNotFound = errors.New("record not found")
	ErrInvalidToken = errors.New("Invalid token")
	ErrUnauthorized = errors.New("Unauthorized")
	ErrForbidden = errors.New("Forbidden")
	ErrInvalidDateFormat = errors.New("Invalid date format. Use DD/MM/YYYY")
)