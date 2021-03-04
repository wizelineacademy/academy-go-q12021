package card

// CardService defines author service behavior.
type CardService interface {
  CreateCard(*Card) (*Card, error)
  ReadCard(id int) (*Card, error)
  ListCards() ([]Card, error)
}

// Service struct handles author business logic tasks.
type Service struct {
  repository CardRepository
}

func (svc *Service) CreateCard(author *Card) (*Card, error) {
  return svc.repository.CreateCard(author)
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
