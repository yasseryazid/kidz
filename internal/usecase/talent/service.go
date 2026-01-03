package talent

import (
	"context"
	"errors"
	"strings"

	domain "github.com/yasseryazid/boilerpart/internal/domain/talent"
)

var ErrInvalidID = errors.New("invalid talent id")

// Service orchestrates talent use cases.
type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List(ctx context.Context) ([]domain.Talent, error) {
	return s.repo.List(ctx)
}

func (s *Service) Create(ctx context.Context, name, title string, skills []string) (domain.Talent, error) {
	talent, err := domain.New(name, title, skills)
	if err != nil {
		return domain.Talent{}, err
	}

	return s.repo.Create(ctx, talent)
}

func (s *Service) GetByID(ctx context.Context, id string) (domain.Talent, error) {
	if strings.TrimSpace(id) == "" {
		return domain.Talent{}, ErrInvalidID
	}

	return s.repo.GetByID(ctx, id)
}
