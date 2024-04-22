package postgres

import (
	"errors"

	"github.com/jackc/pgconn"
)

const (
	UniquenessViolation = "23505"
	LockNotAvailable    = "55P03"
	ForeignKeyViolation = "23503"
)

func IsUniquenessViolation(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == UniquenessViolation {
		return true
	}

	return false
}

func IsForeignKeyViolation(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == ForeignKeyViolation {
		return true
	}

	return false
}

func IsLockNotAvailable(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == LockNotAvailable {
		return true
	}

	return false
}
