package repositories

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/domains/entities"
)

type ITask interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*entities.Task, error)
	AddTask(ctx context.Context, t *entities.Task) error
}

type IUser interface {
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) (entities.User, error)
}
