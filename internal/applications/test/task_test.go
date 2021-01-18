package test

import (
	"context"
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/manabie-com/togo/internal/domains/entities"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAddTask(t *testing.T) {
	app := initTest()
	type fields struct {
	}
	tests := []struct {
		name     string
		fields   fields
		funcMock func()
		wantErr  bool
		want     *entities.Task
	}{
		{
			name:    "insert success",
			fields:  fields{},
			wantErr: false,
			funcMock: func() {
				app.mockTaskDb.On("AddTask", mock.Anything, mock.Anything).Return(nil).Times(1)
			},
			want: &entities.Task{
				ID:          "T",
				Content:     "T",
				UserID:      "T",
				CreatedDate: "T",
			},
		},
		{
			name:    "can't insert",
			fields:  fields{},
			wantErr: true,
			funcMock: func() {
				app.mockTaskDb.On("AddTask", mock.Anything, mock.Anything).Return(fmt.Errorf("test")).Times(1)
			},
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
			tt.funcMock()
			got, err := app.Application.UserAddTask(context.TODO(), &entities.Task{
				ID:          "T",
				Content:     "T",
				UserID:      "T",
				CreatedDate: "T",
			})
			if err != nil && tt.wantErr == false {
				t.Errorf("Have 1 err ")
			}
			assert.Equal(t, got, tt.want)
		})
	}

}
