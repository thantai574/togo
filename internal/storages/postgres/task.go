package postgres

import (
	"context"
	"database/sql"
	"github.com/go-pg/pg/v10"
	"github.com/manabie-com/togo/internal/domains/entities"
)

type taskModel struct {
	db *pg.DB
}

func NewTask(db *pg.DB) *taskModel {
	return &taskModel{db: db}
}

func (t *taskModel) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) (tasks []*entities.Task, err error) {
	return
}

func (t *taskModel) AddTask(ctx context.Context, task *entities.Task) (err error) {
	return
}
