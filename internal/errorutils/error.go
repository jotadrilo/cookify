package errorutils

import (
	"errors"
	"fmt"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrNotFound       = errors.New("not found")
	ErrAlreadyExists  = errors.New("already exists")
	ErrNotCreated     = errors.New("not created")
	ErrBadRequest     = errors.New("bad request")
)

func NewErrNotImplemented(s string) error {
	return fmt.Errorf("%s %w", s, ErrNotImplemented)
}

func NewErrNotFound(s string) error {
	return fmt.Errorf("%s %w", s, ErrNotFound)
}

func NewErrAlreadyExists(s string) error {
	return fmt.Errorf("%s %w", s, ErrAlreadyExists)
}

func NewErrBadRequest(s string) error {
	return fmt.Errorf("%w: %s", ErrBadRequest, s)
}

func NewErrNotCreated(s string) error {
	return fmt.Errorf("%s %w", s, ErrNotCreated)
}
