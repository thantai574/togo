package postgres

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/domains/entities"
	"gorm.io/gorm"
)

type taskModel struct {
	db *gorm.DB
}

func NewTask(db *gorm.DB) *taskModel {
	return &taskModel{db: db}
}

func (t *taskModel) RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) (tasks []*entities.Task, err error) {
	cur, err := t.db.Table("tasks").Where("user_id = ?", userID).Where("created_date = ?", createdDate).Rows()
	if err != nil {
		return
	}
	defer cur.Close()
	for cur.Next() {
		var detail_task = &entities.Task{}
		err = cur.Scan(&detail_task.ID, &detail_task.Content, &detail_task.UserID, &detail_task.CreatedDate)
		if err == nil {
			tasks = append(tasks, detail_task)
		}
	}
	return
}

func (t *taskModel) AddTask(ctx context.Context, task *entities.Task) (err error) {
	err = t.db.Create(task).Error
	return
}
