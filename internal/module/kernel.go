package module

import (
	"github.com/maestre3d/academy-go-q12021/internal/domain"
	"github.com/maestre3d/academy-go-q12021/internal/infrastructure"
	"go.uber.org/fx"
)

// Kernel Shared context main module
var Kernel = fx.Provide(
	infrastructure.NewConfiguration,
	infrastructure.NewZapLogger,
	func() domain.EventBus {
		return nil
	},
)
