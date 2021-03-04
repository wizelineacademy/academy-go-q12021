package cards

import (
  domain "github.com/gbrayhan/academy-go-q12021/domain/card"
)

func toResponseModel(entity *domain.Card) *CardResponse {
  return &CardResponse{
    ID:        entity.ID,
    Name:      entity.Name,
    Type:      entity.Type,
    Level:     entity.Level,
    Race:      entity.Race,
    Attribute: entity.Attribute,
    ATK:       entity.ATK,
    DEF:       entity.DEF,
  }
}
