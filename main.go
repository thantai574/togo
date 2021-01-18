package main

import (
	"github.com/manabie-com/togo/internal/applications"
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

	s.Application = applications.Application{}.
		WithConfig(conf).
		WithLogger(lg).Build()

	http.ListenAndServe(":5050", s)
}
