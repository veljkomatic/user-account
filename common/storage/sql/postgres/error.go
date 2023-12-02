package postgres

import (
	"context"
	"database/sql"
	"os"

	"github.com/pkg/errors"
	"github.com/uptrace/bun/driver/pgdriver"
)

// ErrIntegrityViolation checks if there is a integrity violation, that is if the error is part of the
// Integrity Constraint Violation class of errors.
func ErrIntegrityViolation(err error) bool {
	var pgErr pgdriver.Error
	if errors.As(err, &pgErr) {
		return pgErr.IntegrityViolation()
	}

	// If the error has a cause, check if the cause is an pgdriver.Error
	cause := errors.Cause(err)
	if errors.As(cause, &pgErr) {
		return pgErr.IntegrityViolation()
	}

	return false
}

// ErrNoRows returns if the error's cause is sql.ErrNoRows
// Instead of pg.ErrNoRows bun uses sql.ErrNoRows
func ErrNoRows(err error) bool {
	return errors.Is(errors.Cause(err), sql.ErrNoRows)
}

func HandleErr(err error) error {
	if err == nil {
		return nil
	}
	if ErrNoRows(err) {
		return errors.New("NotFound")
	}
	if errors.Is(err, context.Canceled) {
		return errors.New("Canceled context")
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return errors.New("Deadline exceeded")
	}
	// "i/o timeout" is returned by pgx when the query deadline is exceeded.
	if errors.Is(err, os.ErrDeadlineExceeded) {
		return errors.New("OS Deadline exceeded")
	}
	if ErrIntegrityViolation(err) {
		return errors.New("Integrity violation")
	}
	return errors.New("Unknown error")
}

func HandleAndWrapErr(err error, message string) error {
	return errors.Wrap(HandleErr(err), message)
}
