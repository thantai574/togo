package sqllite

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/storages"
)

type userModel struct {
	DB *sql.DB
}

// ValidateUser returns tasks if match userID AND password
func (l *userModel) ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool {
	stmt := `SELECT id FROM users WHERE id = ? AND password = ?`
	row := l.DB.QueryRowContext(ctx, stmt, userID, pwd)
	u := &storages.User{}
	err := row.Scan(&u.ID)
	if err != nil {
		return false
	}

	return true
}
