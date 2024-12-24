package api

import (
	"aari/web_api/httputil"
	"aari/web_api/httputil/httperror"
	"net/http"

	"github.com/uptrace/bunrouter"
)

var version = "v1"

type API struct {
	Version     string
	Env         string
	Healthcheck bunrouter.HandlerFunc
	Users       UserHandler
	Positions   PositionHandler
}

func (api *API) healthcheck(w http.ResponseWriter, r bunrouter.Request) error {
	data := map[string]string{
		"status":  "ok",
		"env":     api.Env,
		"version": version,
	}
	if err := httputil.JsonResponse(w, http.StatusOK, data); err != nil {
		return httperror.ErrInternal
	}
	return nil
}

func NewAPI(app *App, env string) *API {

	return &API{
		Version:   version,
		Env:       env,
		Users:     UserHandler{app},
		Positions: PositionHandler{app},
	}

}
