package api

import (
	"aari/web_api/internal/db"
	"aari/web_api/internal/models"
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/uptrace/bun"
	"github.com/uptrace/bunrouter"
	"github.com/urfave/cli/v2"
)

type appCtxKey struct{}

func AppFromContext(ctx context.Context) *App {
	return ctx.Value(appCtxKey{}).(*App)
}

func ContextWithApp(ctx context.Context, app *App) context.Context {
	ctx = context.WithValue(ctx, appCtxKey{}, app)
	return ctx
}

type App struct {
	ctx         context.Context
	cfg         *Config
	router      *bunrouter.Router
	apiRouter   *bunrouter.Group
	dbOnce      sync.Once
	db          *bun.DB
	onStop      appHooks
	onAfterStop appHooks
	store       *models.Store
}

func NewApp(ctx context.Context, cfg *Config) *App {

	app := &App{
		cfg: cfg,
	}

	app.ctx = ContextWithApp(ctx, app)
	app.InitRouter()
	app.store = models.InitStore(app.DB())
	return app
}
func StartCLI(c *cli.Context) (context.Context, *App, error) {
	return Start(c.Context, c.Command.Name, c.String("env"))
}

func Start(ctx context.Context, service, envName string) (context.Context, *App, error) {
	cfg, err := ReadConfig(FS(), service, envName)
	if err != nil {
		return nil, nil, err
	}
	return StartFromConfig(ctx, cfg)
}

func StartFromConfig(ctx context.Context, cfg *Config) (context.Context, *App, error) {
	app := NewApp(ctx, cfg)
	if err := onStart.Run(ctx, app); err != nil {
		return nil, nil, err
	}
	return app.Context(), app, nil
}
func (app *App) Store() *models.Store {
	return app.store
}

func (app *App) IsDebug() bool {
	return app.cfg.Debug
}

func (app *App) Context() context.Context {
	return app.ctx
}

func (app *App) Router() *bunrouter.Router {
	return app.router
}

func (app *App) APIRouter() *bunrouter.Group {
	return app.apiRouter
}
func (app *App) Stop() {
	_ = app.onStop.Run(app.ctx, app)
	_ = app.onAfterStop.Run(app.ctx, app)
}

func (app *App) OnStop(name string, fn HookFunc) {
	app.onStop.Add(newHook(name, fn))
}

func (app *App) OnAfterStop(name string, fn HookFunc) {
	app.onAfterStop.Add(newHook(name, fn))
}

func (app *App) DB() *bun.DB {
	app.dbOnce.Do(
		func() {
			db, err := db.New(
				app.cfg.PGX.DSN,
				app.cfg.PGX.maxOpenConns,
				app.cfg.PGX.maxIdleConns,
				app.cfg.PGX.maxIdleTime,
			)
			if err != nil {
				log.Fatal(err)
				return
			}
			app.db = db
		},
	)
	return app.db
}

func WaitExitSignal() os.Signal {
	ch := make(chan os.Signal, 3)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	return <-ch
}
