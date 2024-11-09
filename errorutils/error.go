package errorutils

import (
	"errors"
	"fmt"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrNotFound       = errors.New("not found")
	ErrBadRequest     = errors.New("bad request")
)

func NewErrNotImplemented(s string) error {
	return fmt.Errorf("%s %w", s, ErrNotImplemented)
}

func NewErrNotFound(s string) error {
	return fmt.Errorf("%s %w", s, ErrNotFound)
}

func NewErrBadRequest(s string) error {
	return fmt.Errorf("%w: %s", ErrBadRequest, s)
}
