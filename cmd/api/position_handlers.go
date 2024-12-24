package api

import (
	"aari/web_api/httputil"
	"aari/web_api/httputil/httperror"
	"aari/web_api/internal/models"
	"net/http"
	"time"

	"github.com/uptrace/bunrouter"
)

type PositionHandler struct {
	app *App
}

type positionCreatePayload struct {
	Ticker          string    `json:"ticker"`
	AssetType       string    `json:"asset_type"`
	Bias            string    `json:"bias"`
	Justification   string    `json:"justification"`
	Expiration      time.Time `json:"expiration"`
	CapitalInvested float64   `json:"capital_invested"`
}

func (ph *PositionHandler) Create(w http.ResponseWriter, r bunrouter.Request) error {
	var payload positionCreatePayload
	if err := httputil.ReadJSON(w, r.Request, &payload); err != nil {
		return httperror.BadRequest("json_syntax", err.Error())
	}

	position := &models.PositionCreate{
		Ticker:          payload.Ticker,
		AssetType:       payload.AssetType,
		Bias:            payload.Bias,
		Justification:   payload.Justification,
		Expiration:      payload.Expiration,
		CapitalInvested: payload.CapitalInvested,
	}

	err := ph.app.Store().Positions.Create(r.Context(), position)
	if err != nil {
		return httperror.From(err)
	}
	return nil
}
