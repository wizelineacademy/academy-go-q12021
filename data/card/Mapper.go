package card

import (
  domain "github.com/gbrayhan/academy-go-q12021/domain/card"
)

func ToDataModel(entity *domain.Card) *Card {
  return &Card{
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

func ToDomainModel(entity *Card) *domain.Card {
  return &domain.Card{
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
