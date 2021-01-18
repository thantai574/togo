package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/manabie-com/togo/internal/utils/configs"
)

func NewDB(conf *configs.Config) (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		User:     conf.DBPostgres.UserName,
		Password: conf.DBPostgres.Password,
		Addr:     conf.DBPostgres.Addr,
	})

	return
}
