package test

import (
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/utils/configs"
	"gorm.io/gorm"
)

func initTestDB() (*gorm.DB, error) {
	conf, err := configs.LoadConfig("../../../../", "config_test.json")
	if err != nil {
		panic(err)
	}
	return postgres.NewDB(conf)
}
