package card

// Card struct defines the database model for a card.
type Card struct {
  ID        int
  Name      string
  Type      string
  Level     int
  Race      string
  Attribute string
  ATK       int
  DEF       int
}
