package postgres

import (
	"context"
	"database/sql"
	"github.com/go-pg/pg/v10"
	"github.com/manabie-com/togo/internal/domains/entities"
)

type userModel struct {
	db *pg.DB
}

func NewUserModel(db *pg.DB) *userModel {
	return &userModel{db: db}
}

func (u userModel) ValidateUser(ctx context.Context, userID, pwd sql.NullString) (user entities.User, err error) {
	return
}
