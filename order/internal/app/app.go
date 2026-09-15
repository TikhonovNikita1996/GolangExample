package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/pkg/errors"

	"github.com/TikhonovNikita1996/Go-microservises-project/order/internal/config"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/closer"
	"github.com/TikhonovNikita1996/Go-microservises-project/platform/pkg/logger"
	orderv1 "github.com/TikhonovNikita1996/Go-microservises-project/shared/pkg/openapi/order/v1"
)

type App struct {
	diContainer *diContainer
	orderServer *orderv1.Server
}

func New(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error { return a.runHTTPServer(ctx) }

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initDI,
		a.initLogger,
		a.initCloser,
		a.runConsumer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initDI(_ context.Context) error {
	a.diContainer = NewDiContainer()
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	return logger.Init(
		config.AppConfig().Logger.Level(),
		config.AppConfig().Logger.AsJson(),
	)
}

func (a *App) initCloser(_ context.Context) error {
	closer.SetLogger(logger.Logger())
	return nil
}

func (a *App) runHTTPServer(ctx context.Context) error {
	r := chi.NewRouter()

	orderServer, err := orderv1.NewServer(a.diContainer.OrderV1API(ctx))
	if err != nil {
		log.Fatalf("Ошибка создания сервера OpenAPI: %v", err)
	}

	server := &http.Server{
		Addr:              net.JoinHostPort(config.AppConfig().OrderHTTPConfig.Host(), config.AppConfig().OrderHTTPConfig.Port()),
		Handler:           r,
		ReadHeaderTimeout: config.AppConfig().OrderHTTPConfig.ReadTimeout(),
	}

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Mount("/", orderServer)

	logger.Info(ctx, fmt.Sprintf("🚀 HTTP-сервер запущен на порту %s\n", config.AppConfig().OrderHTTPConfig.Port()))
	ok := server.ListenAndServe()
	if ok != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Errorf("❌ Ошибка запуска сервера: %v\n", ok))
	}

	closer.AddNamed("HTTP server", func(ctx context.Context) error {
		return server.Shutdown(ctx)
	})

	return nil
}

func (a *App) runConsumer(ctx context.Context) error {
	logger.Info(ctx, "🚀 ShipAssembled Kafka consumer running")

	go func() {
		if err := a.diContainer.OrderConsumerService().RunConsumer(ctx); err != nil {
			logger.Error(ctx, "❌ Kafka consumer error: "+err.Error())
		}
	}()

	return nil
}
