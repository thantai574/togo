package test

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"testing"
)

func TestUser(t *testing.T) {
	db, _ := initTestDB()
	userDB := postgres.NewUserModel(db)
	tests := []struct {
		name     string
		funcMock func()
		wantErr  bool
		want     struct {
			id string
		}
	}{
		{
			name:    "check user",
			wantErr: false,
			want: struct {
				id string
			}{
				id: "firstUser",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := userDB.ValidateUser(context.TODO(), sql.NullString{
				String: "firstUser",
				Valid:  true,
			}, sql.NullString{
				String: "example",
				Valid:  true,
			})
			if err != nil && tt.wantErr == false {
				t.Errorf("Have 1 err AddTask")
			}

			if u.ID != tt.want.id {
				t.Errorf("Have 1 err user")
			}

		})
	}

}
