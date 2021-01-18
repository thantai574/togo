package test

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/domains/entities"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"testing"
	"time"
)

func TestTask(t *testing.T) {
	db, _ := initTestDB()
	taskDb := postgres.NewTask(db)
	defer db.Exec("DELETE FROM tasks")
	now := time.Now()
	tests := []struct {
		name     string
		funcMock func()
		wantErr  bool
		want     struct {
			countTotal int
		}
	}{
		{
			name:    "insert success & get list task",
			wantErr: false,
			want: struct {
				countTotal int
			}{
				countTotal: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := taskDb.AddTask(context.TODO(), &entities.Task{
				ID:          uuid.New().String(),
				Content:     "T",
				UserID:      "firstUser",
				CreatedDate: now.Format("2006-01-02"),
			})
			if err != nil && tt.wantErr == false {
				t.Errorf("Have 1 err AddTask")
			}
			tasks, err := taskDb.RetrieveTasks(context.TODO(), sql.NullString{
				String: "firstUser",
				Valid:  true,
			}, sql.NullString{
				String: now.Format("2006-01-02"),
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
