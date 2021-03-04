package card

// CardRepository provides an abstraction on top of the card data source
type CardRepository interface {
  CreateCard(*Card) (*Card, error)
  ReadCard(int) (*Card, error)
  ListCards() ([]Card, error)
}
