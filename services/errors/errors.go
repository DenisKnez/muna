package errors

import (
	"errors"
)

var (
	ErrStringToLong = errors.New("The string is not allowed to be longer than 100 characters")

	ErrFailedRegexCompilation = errors.New("Regex failed to compile")
)
