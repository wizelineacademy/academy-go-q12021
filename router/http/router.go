package http

import (
  "net/http"

  "github.com/gbrayhan/academy-go-q12021/domain/card"

  "github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"

  cardsRoutes "github.com/gbrayhan/academy-go-q12021/router/http/cards"
  "github.com/gbrayhan/academy-go-q12021/router/http/errors"
  healthRoutes "github.com/gbrayhan/academy-go-q12021/router/http/health"
)

// NewHTTPHandler returns the HTTP requests handler
func NewHTTPHandler(cardSvc card.CardService) http.Handler {
  router := gin.Default()

  config := cors.DefaultConfig()
  config.AllowAllOrigins = true
  config.AddAllowHeaders("Authorization")
  router.Use(cors.New(config))

  router.Use(errors.Handler)

  healthGroup := router.Group("/health")
  healthRoutes.NewRoutesFactory()(healthGroup)

  api := router.Group("/api")

  cardsGroup := api.Group("/cards")
  cardsRoutes.NewRoutesFactory(cardsGroup)(cardSvc)
  // booksRoutes.NewRoutesFactory(cardsGroup)(booksSvc)
  return router
}
