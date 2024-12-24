package api

import (
	"aari/web_api/httputil"
	"aari/web_api/httputil/httperror"
	"aari/web_api/internal/models"
	"net/http"

	"github.com/uptrace/bunrouter"
)

type UserHandler struct {
	app *App
}

type registerUserPayload struct {
	username string `json:"username"`
	email    string `json:"email"`
	password string `json:"password"`
}

func (uh *UserHandler) Register(w http.ResponseWriter, r bunrouter.Request) error {
	var payload registerUserPayload
	if err := httputil.ReadJSON(w, r.Request, &payload); err != nil {
		return httperror.BadRequest("json_syntax", err.Error())
	}

	user := &models.UserCreate{
		Username: payload.username,
		Email:    payload.email,
		Password: payload.password,
	}
	err := uh.app.Store().Users.Create(r.Context(), user)
	if err != nil {
		return httperror.From(err)
	}
	return nil
}
