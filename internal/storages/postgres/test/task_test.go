package test

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/domains/entities"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"testing"
)

func TestTask(t *testing.T) {
	db, _ := initTestDB()
	taskDb := postgres.NewTask(db)
	defer db.Exec("DELETE FROM task")
	tests := []struct {
		name     string
		funcMock func()
		wantErr  bool
		want     struct {
			insert     *entities.Task
			countTotal int
		}
	}{
		{
			name:    "insert success & get list task",
			wantErr: false,
			want: struct {
				insert     *entities.Task
				countTotal int
			}{
				insert: &entities.Task{
					ID:          "T",
					Content:     "T",
					UserID:      "T",
					CreatedDate: "2020-01-18",
				},
				countTotal: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := taskDb.AddTask(context.TODO(), &entities.Task{
				ID:          "T",
				Content:     "T",
				UserID:      "T",
				CreatedDate: "2020-01-18",
			})
			if err != nil && tt.wantErr == false {
				t.Errorf("Have 1 err AddTask")
			}
			tasks, err := taskDb.RetrieveTasks(context.TODO(), sql.NullString{
				String: "T",
				Valid:  true,
			}, sql.NullString{
				String: "2020-01-18",
				Valid:  true,
			})
			if err != nil && tt.wantErr == false {
				t.Errorf("Have 1 err RetrieveTasks ")
			}

			if len(tasks) != tt.want.countTotal {
				t.Errorf("RetrieveTasks err ")
			}

		})
	}

}
