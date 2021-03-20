package usecase

import (
	"fmt"
	"errors"
	"github.com/halarcon-wizeline/academy-go-q12021/domain"
)

// UseCase struct
type UseCase struct {
	domain domain.Pokemon
}

// New UseCase
func New(domain domain.Pokemon) *UseCase {
	return &UseCase{domain}
}
