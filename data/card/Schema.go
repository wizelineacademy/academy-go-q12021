package card

// struct defines the database model for a Author.
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
