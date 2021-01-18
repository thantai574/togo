package main

import (
	"github.com/manabie-com/togo/internal/applications"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/utils/configs"
	"github.com/manabie-com/togo/internal/utils/logger"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	s := &services.ToDoService{}

	conf, err := configs.LoadConfig("./", "config.json")
	if err != nil {
		panic(err)
	}
	s.Config = conf

	lg, err := logger.NewLogger(conf.Env)
	if err != nil {
		panic(err)
	}
	app := &applications.Application{}
	db, _ := postgres.NewDB(conf)
	s.Application = app.
		WithConfig(conf).
		WithLogger(lg).
		WithTaskStorage(postgres.NewTask(db)).
		WithUserStorage(postgres.NewUserModel(db)).Build()

	http.ListenAndServe(":5050", s)
}
