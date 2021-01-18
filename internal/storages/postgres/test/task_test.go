package test

import (
	"context"
	"github.com/manabie-com/togo/internal/domains/entities"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"testing"
)

func TestAddTask(t *testing.T) {
	db := initTestDB()
	taskDb := postgres.NewTask(db)
	tests := []struct {
		name     string
		funcMock func()
		wantErr  bool
		want     *entities.Task
	}{
		{
			name:    "insert success",
			wantErr: false,
			want: &entities.Task{
				ID:          "T",
				Content:     "T",
				UserID:      "T",
				CreatedDate: "T",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := taskDb.AddTask(context.TODO(), &entities.Task{
				ID:          "T",
				Content:     "T",
				UserID:      "T",
				CreatedDate: "T",
			})
			if err != nil && tt.wantErr == false {
				t.Errorf("Have 1 err ")
			}
		})
	}

}

func TestGetTask(t *testing.T) {

}
