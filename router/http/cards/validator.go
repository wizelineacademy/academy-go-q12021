package cards

import (
  "github.com/gin-gonic/gin"

  "github.com/gbrayhan/academy-go-q12021/domain/card"
)

// AuthorValidator is a struct used to validate the JSON payload representing an author.
type CardValidator struct {
  Name string `binding:"required" json:"name"`
  Type string `binding:"required" json:"type"`
}

// Bind validates the JSON payload and returns a Author instance corresponding to the payload.
func Bind(c *gin.Context) (*card.Card, error) {
  var json CardValidator
  if err := c.ShouldBindJSON(&json); err != nil {
    return nil, err
  }

  dataCard := &card.Card{
    Name: json.Name,
    Type: json.Type,
  }

  return dataCard, nil
}
