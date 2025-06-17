package customerror

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrUniqueViolation     = errors.New("unique constraint violation")
	ErrForeignKeyViolation = errors.New("foreign key constraint violation")
	ErrCheckViolation      = errors.New("check constraint violation")
	ErrNotNullViolation    = errors.New("not null constraint violation")
	ErrDatabaseError       = errors.New("database error")
)

const (
	UniqueViolation     = "23505"
	ForeignKeyViolation = "23503"
	CheckViolation      = "23514"
	NotNullViolation    = "23502"
)

func HandlePostgresError(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case UniqueViolation:
			column := extractColumnFromConstraint(pgErr.ConstraintName)
			return fmt.Errorf("%w: %s already exists", ErrUniqueViolation, column)
		case ForeignKeyViolation:
			column := extractColumnFromConstraint(pgErr.ConstraintName)
			return fmt.Errorf("%w: invalid value for column: %s", ErrForeignKeyViolation, column)
		case CheckViolation:
			return fmt.Errorf("%w: %s", ErrCheckViolation, pgErr.ConstraintName)
		case NotNullViolation:
			return fmt.Errorf("%w: column %s cannot be null", ErrNotNullViolation, pgErr.ColumnName)
		default:
			return fmt.Errorf("%w: %v", ErrDatabaseError, err)
		}
	}
	return err
}

func extractColumnFromConstraint(constraint string) string {
	parts := strings.Split(constraint, "_")
	if len(parts) >= 3 {
		return parts[len(parts)-1]
	}
	return "field"
}
