package applications

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/domains/entities"
	"github.com/manabie-com/togo/internal/errors"
)

func (app *Application) UserAddTask(ctx context.Context, createTask *entities.Task) (task *entities.Task, err error) {
	err = app.iTask.AddTask(ctx, createTask)

	if err != nil {
		err = errors.NewErrorMsg(errors.ListTasksFailed, err.Error(), errors.AddTaskFailed.ExternalErrString())
	}

	return createTask, err
}

func (app *Application) UserGetTasks(ctx context.Context, userID, createDate sql.NullString) (tasks []*entities.Task, err error) {
	tasks, err = app.iTask.RetrieveTasks(ctx, userID, createDate)

	if err != nil {
		err = errors.NewErrorMsg(errors.ListTasksFailed, err.Error(), errors.ListTasksFailed.ExternalErrString())
	}

	return
}
