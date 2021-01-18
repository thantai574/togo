package aggregates

import "github.com/manabie-com/togo/internal/domains/entities"

type UserProfile struct {
	entities.User
	Token string `json:"token"`
}
