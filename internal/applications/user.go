package applications

import (
	"context"
	"database/sql"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/domains/aggregates"
	"go.uber.org/zap/zapcore"
	"time"
)

func (app *Application) UserLogin(ctx context.Context, userID, password sql.NullString) (user aggregates.UserProfile, err error) {
	entity_user, err := app.iUser.ValidateUser(ctx, userID, password)
	if err != nil {
		app.logger.With(zapcore.Field{
			Key:       "err",
			Type:      zapcore.ReflectType,
			Interface: err,
		}).Info("ValidateUser err ")
		return
	}
	user.User = entity_user

	token, err := app.UserCreateToken(entity_user.ID)

	if err != nil {
		app.logger.With(zapcore.Field{
			Key:       "err",
			Type:      zapcore.ReflectType,
			Interface: err,
		}).Info("UserLogin get token err ")
		return
	}

	user.Token = token
	return

}

func (app *Application) UserCreateToken(id string) (token string, err error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err = at.SignedString([]byte(app.config.JwtKey))

	return
}
