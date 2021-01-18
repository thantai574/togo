package test

import (
	"github.com/go-pg/pg/v10"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/utils/configs"
)

func initTestDB() *pg.DB {
	conf, err := configs.LoadConfig("../../../../", "config_test.json")
	if err != nil {
		panic(err)
	}
	return postgres.NewDB(conf)
}
