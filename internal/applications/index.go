package applications

import (
	"github.com/manabie-com/togo/internal/domains/repositories"
	"github.com/manabie-com/togo/internal/utils/configs"
	"go.uber.org/zap"
)

type Application struct {
	iTask  repositories.ITask
	iUser  repositories.IUser
	logger *zap.Logger
	config *configs.Config
}

func (app *Application) WithTaskStorage(iTask repositories.ITask) *Application {
	app.iTask = iTask
	return app
}

func (app *Application) WithUserStorage(iUser repositories.IUser) *Application {
	app.iUser = iUser
	return app
}

func (app *Application) WithConfig(config *configs.Config) *Application {
	app.config = config
	return app
}

func (app *Application) WithLogger(logger *zap.Logger) *Application {
	app.logger = logger
	return app
}

func (app *Application) Build() *Application {
	return app
}
