package postgres

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/domains/entities"
	"gorm.io/gorm"
)

type userModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *userModel {
	return &userModel{db: db}
}

func (u userModel) ValidateUser(ctx context.Context, userID, pwd sql.NullString) (user entities.User, err error) {
	err = u.db.Table("users").Where("id = ?", userID).Where("password = ?", pwd).First(&user).Error
	if err != nil {
		return
	}
	return
}
