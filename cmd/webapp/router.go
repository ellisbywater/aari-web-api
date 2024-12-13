package webapp

import (
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/bunrouterotel"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func (app *App) InitRouter() {
	app.router = bunrouter.New(
		bunrouter.WithMiddleware(reqlog.NewMiddleware(
			reqlog.WithEnabled(app.IsDebug()),
			reqlog.FromEnv(""),
		)),
		bunrouter.WithMiddleware(bunrouterotel.NewMiddleware()),
	)

	app.apiRouter = app.router.NewGroup("/api",
		bunrouter.WithMiddleware(corsMiddleware),
		bunrouter.WithMiddleware(errorHandler),
	)
}
