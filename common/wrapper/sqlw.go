package sqlw

import (
	"context"
	"database/sql"
)

type SQLWrapper interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}
