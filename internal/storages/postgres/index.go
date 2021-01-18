package postgres

import (
	"fmt"
	"github.com/manabie-com/togo/internal/utils/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(conf *configs.Config) (db *gorm.DB, err error) {

	connectString := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		conf.DBPostgres.Addr,
		conf.DBPostgres.UserName,
		conf.DBPostgres.Password,
		conf.DBPostgres.DB,
		conf.DBPostgres.Port,
		conf.DBPostgres.SSLMode,
		conf.DBPostgres.TimeZone,
	)
	db, err = gorm.Open(postgres.Open(connectString), &gorm.Config{})

	return
}
