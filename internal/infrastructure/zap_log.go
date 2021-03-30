package infrastructure

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// NewZapLogger creates a new Uber's Zap logger depending on the development stage
func NewZapLogger(lc fx.Lifecycle, cfg Configuration) (logger *zap.Logger, err error) {
	if cfg.IsProd() {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	var stdClose func()
	if logger != nil {
		stdClose = zap.RedirectStdLog(logger)
	}

	lc.Append(fx.Hook{
		OnStart: nil,
		OnStop: func(ctx context.Context) error {
			if err := logger.Sync(); logger != nil && err != nil {
				return err
			}
			if stdClose != nil {
				stdClose()
			}
			return nil
		},
	})
	return
}
