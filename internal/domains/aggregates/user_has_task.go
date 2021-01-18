package aggregates

import "github.com/manabie-com/togo/internal/domains/entities"

type UserHasTask struct {
	entities.User
	Tasks []entities.Task
}
