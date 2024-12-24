package api

import (
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/bunrouterotel"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func (app *App) InitRouter() {
	api := NewAPI(app, "dev")
	app.router = bunrouter.New(
		bunrouter.WithMiddleware(reqlog.NewMiddleware(
			reqlog.WithEnabled(app.IsDebug()),
			reqlog.FromEnv(""),
		)),
		bunrouter.WithMiddleware(bunrouterotel.NewMiddleware()),
	)

	app.apiRouter = app.router.NewGroup("/v1/api",
		bunrouter.WithMiddleware(corsMiddleware),
		bunrouter.WithMiddleware(errorHandler),
	)
	app.apiRouter.GET("/healthcheck", api.healthcheck)
}
