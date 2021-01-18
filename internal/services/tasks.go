package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/manabie-com/togo/internal/applications"
	"github.com/manabie-com/togo/internal/domains/entities"
	"github.com/manabie-com/togo/internal/errors"
	"github.com/manabie-com/togo/internal/utils/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// ToDoService implement HTTP server
type ToDoService struct {
	Logger *zap.Logger
	Config *configs.Config
	*applications.Application
}

func (s *ToDoService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	defer func() {
		if r := recover(); r != nil {
			s.Logger.With(zapcore.Field{
				Key:       "recovery_data",
				Type:      zapcore.ReflectType,
				Interface: r,
			}).Info("recovery panic ")
		}
	}()
	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	switch req.URL.Path {
	case "/login":
		s.getAuthToken(resp, req)
		return
	case "/tasks":
		var ok bool
		req, ok = s.validToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch req.Method {
		case http.MethodGet:
			s.listTasks(resp, req)
		case http.MethodPost:
			s.addTask(resp, req)
		}
		return
	}
}

func customError(resp http.ResponseWriter, err error) {
	if e, ok := err.(errors.ErrorApplication); ok {
		json.NewEncoder(resp).Encode(e)
	} else {
		json.NewEncoder(resp).Encode(errors.ErrorApplication{
			Code:        400,
			InternalMsg: err.Error(),
			ExternalMsg: err.Error(),
		})
	}
}

func value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

func customSuccess(resp http.ResponseWriter, data interface{}) {
	json.NewEncoder(resp).Encode(map[string]interface{}{
		"data": data,
	})
}

func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := value(req, "user_id")
	u, err := s.Application.UserLogin(req.Context(), id, value(req, "password"))
	if err != nil {
		customError(resp, err)
	}
	resp.Header().Set("Content-Type", "application/json")
	customSuccess(resp, u)
}

func (s *ToDoService) listTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	tasks, err := s.Application.UserGetTasks(
		req.Context(),
		sql.NullString{
			String: id,
			Valid:  true,
		},
		value(req, "created_date"),
	)

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		customError(resp, err)
		return
	}

	customSuccess(resp, tasks)
}

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	t := &entities.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()

	if err != nil {
		customError(resp, err)
		return
	}

	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	resp.Header().Set("Content-Type", "application/json")

	t, err = s.UserAddTask(req.Context(), t)

	if err != nil {
		customError(resp, err)
		return
	}

	customSuccess(resp, t)
}

func (s *ToDoService) validToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.Config.JwtKey), nil
	})
	if err != nil {
		log.Println(err)
		return req, false
	}

	if !t.Valid {
		return req, false
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), id))
	return req, true
}

type userAuthKey int8

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
