package errors

// Refactor this package to be sugared wrapper around pkg/errors
// It should have error type for each error that can occur in the application
// For example:
// ErrTypeNotFound
// ErrTypeInvalidInput
// ErrTypeConflict

import "github.com/pkg/errors"

var ErrNotFound = errors.New("not found")
