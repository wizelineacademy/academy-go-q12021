package httputil

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	muxhandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/maestre3d/academy-go-q12021/internal/infrastructure"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// StartServer bootstraps a pre-configured HTTP Server
func StartServer(lc fx.Lifecycle, cfg infrastructure.Configuration, logger *zap.Logger, mux *mux.Router,
	gw ControllersGatewayFx) {
	prefix := NewVersioning(cfg)
	routeControllers(prefix, mux, gw)
	logServer(cfg, logger, prefix)
	startServer(lc, cfg, mux)
}

func routeControllers(prefix string, serveMux *mux.Router, gateway ControllersGatewayFx) {
	public := serveMux.PathPrefix(prefix).Subrouter()
	for _, c := range gateway.Controllers {
		c.MapRoutes(public)
	}
}

func logServer(cfg infrastructure.Configuration, logger *zap.Logger, prefix string) {
	logger.With(
		zap.Namespace("metadata"),
		zap.String("stage", cfg.Stage),
		zap.String("version", cfg.Version),
		zap.String("address", cfg.HTTPAddress),
		zap.Int("port", cfg.HTTPPort),
		zap.String("prefix", prefix),
	).Info("starting http server at prefix: " + prefix)
}

func startServer(lc fx.Lifecycle, cfg infrastructure.Configuration, mux *mux.Router) {
	addr := cfg.HTTPAddress + ":" + strconv.Itoa(cfg.HTTPPort)
	server := http.Server{
		Addr:              addr,
		Handler:           muxhandlers.CombinedLoggingHandler(os.Stdout, mux),
		TLSConfig:         nil, // use TLS
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) (err error) {
			go func() {
				err = server.ListenAndServe()
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
