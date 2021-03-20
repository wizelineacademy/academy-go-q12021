package controller

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "strconv"
  "github.com/sirupsen/logrus"
  "github.com/unrolled/render"
  "github.com/halarcon-wizeline/academy-go-q12021/domain"
  "github.com/halarcon-wizeline/academy-go-q12021/infrastructure/datastore"
  "github.com/halarcon-wizeline/academy-go-q12021/infrastructure/externalapi"
)

// UseCase interface
type UseCase interface {
}

// Controller struct
type Controller struct {
  useCase UseCase
  logger  *logrus.Logger
  render  *render.Render
}

// New returns a controller
func New(
  u UseCase,
  logger *logrus.Logger,
  r *render.Render,
) *Controller {
  return &Controller{u, logger, r}
}

// GetExternalPokemons logic
func (c *Controller) GetExternalPokemons(w http.ResponseWriter, r *http.Request) {
  fmt.Println("[controller] GetExternalPokemons")

  logger := c.logger.WithFields(logrus.Fields{"func": "Get Pokemons Api", })
  logger.Debug("in")

  pokemons, err := externalapi.GetPokemons()
  if err != nil {
    logger.WithError(err).Error("Getting external api")
    c.render.Text(w, 200, "Can't find api")
    return
  }

  file, err := datastore.CreatePokemonDB("./infrastructure/datastore/pokemons_api.csv", pokemons)
  if err != nil {
    logger.WithError(err).Error("Exporting csv file")
    c.render.Text(w, 200, "Export failed")
    return
  }
  fmt.Println(file)

  c.render.Text(w, 200, file + "created")
}

// GetLocalPokemons logic
func (c *Controller) GetLocalPokemons(w http.ResponseWriter, r *http.Request) {
  fmt.Println("[controller] GetLocalPokemons")

  logger := c.logger.WithFields(logrus.Fields{"func": "Get Local Pokemons", })
  logger.Debug("in")

  pokemons, err := datastore.GetPokemonDB()
  if err != nil {
    logger.WithError(err).Error("Getting local pokemons")
    c.render.Text(w, 200, "Pokemons db not found")
    return 
  }

  json.NewEncoder(w).Encode(pokemons)
}

// GetLocalPokemon logic
func (c *Controller) GetLocalPokemon(w http.ResponseWriter, r *http.Request) {
  fmt.Println("[controller] GetLocalPokemon")

  logger := c.logger.WithFields(logrus.Fields{"func": "Get Local Pokemon", })
  logger.Debug("in")

  pokemons, err := datastore.GetPokemonDB()
  if err != nil {
    logger.WithError(err).Error("Getting local pokemon")
    c.render.Text(w, 200, "Pokemons db not found")
    return 
  }

  params := mux.Vars(r)
  for _, pokemon := range pokemons {
    val, err := strconv.Atoi(params["id"])
    if err != nil {
      logger.WithError(err).Error("Couldn't convert value")
      c.render.Text(w, 200, "Couldn't convert value")
      return 
    }

    if pokemon.ID == val {
      json.NewEncoder(w).Encode(pokemon)
      return
    }
  }
  json.NewEncoder(w).Encode(&domain.Pokemon{})
}
