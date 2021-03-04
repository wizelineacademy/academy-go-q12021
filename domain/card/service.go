package card

// CardService defines card service behavior.
type CardService interface {
  CreateCard(*Card) (*Card, error)
  ReadCard(id int) (*Card, error)
  ListCards() ([]Card, error)
}

// Service struct handles card business logic tasks.
type Service struct {
  repository CardRepository
}

func (svc *Service) CreateCard(card *Card) (*Card, error) {
  return svc.repository.CreateCard(card)
}

func (svc *Service) ReadCard(id int) (*Card, error) {
  return svc.repository.ReadCard(id)
}

func (svc *Service) ListCards() ([]Card, error) {
  return svc.repository.ListCards()
}

// NewService creates a new service struct
func NewService(repository CardRepository) *Service {
  return &Service{repository: repository}
}
