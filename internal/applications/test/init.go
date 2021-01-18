package test

import (
	"github.com/manabie-com/togo/internal/applications"
	"github.com/manabie-com/togo/internal/domains/repositories/mocks"
	"github.com/manabie-com/togo/internal/utils/configs"
	"github.com/manabie-com/togo/internal/utils/logger"
)

type serviceTest struct {
	mockTaskDb  *mocks.ITask
	mockUserDB  *mocks.IUser
	Application *applications.Application
}

func initTest() (appTest serviceTest) {
	conf, err := configs.LoadConfig("../../../", "config_test.json")
	lg, err := logger.NewLogger(conf.Env)
	if err != nil {
		panic(err)
	}
	mockTaskDb := &mocks.ITask{}
	mockUserDB := &mocks.IUser{}
	app := &applications.Application{}

	return serviceTest{mockTaskDb: mockTaskDb, mockUserDB: mockUserDB, Application: app.
		WithConfig(conf).
		WithLogger(lg).
		WithTaskStorage(mockTaskDb).
		WithUserStorage(mockUserDB).Build()}
}
