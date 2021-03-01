package httputil

import (
	"strconv"

	"github.com/maestre3d/academy-go-q12021/internal/infrastructure"
)

// NewVersioning generates a valid URI with given version
func NewVersioning(cfg infrastructure.Configuration) string {
	if major := cfg.MajorVersion(); major > 0 {
		return "/v" + strconv.Itoa(major)
	}
	return "/" + cfg.ReleaseStage()
}
