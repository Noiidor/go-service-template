package service

import (
	"context"

	"github.com/Noiidor/go-service-template/internal/domain"
	"github.com/Noiidor/go-service-template/internal/repos"
)

type WizardsService interface {
	GetByID(ctx context.Context, id uint32) (*domain.Wizard, error)
	GetAll(ctx context.Context) ([]*domain.Wizard, error)
	Create(ctx context.Context, wizard *domain.Wizard) error
	Update(ctx context.Context, wizard *domain.Wizard) error
	Delete(ctx context.Context, id uint32) error
}

type wizardsService struct {
	repo repos.WizardsRepo
}

func NewWizardsService() *wizardsService {
	return &wizardsService{}
}

func (s *wizardsService) GetByID(ctx context.Context, id uint32) (*domain.Wizard, error) {
	return nil, nil
}

func (s *wizardsService) GetAll(ctx context.Context) ([]*domain.Wizard, error) {
	return nil, nil
}

func (s *wizardsService) Create(ctx context.Context, wizard *domain.Wizard) error {
	return nil
}

func (s *wizardsService) Update(ctx context.Context, wizard *domain.Wizard) error {
	return nil
}

func (s *wizardsService) Delete(ctx context.Context, id uint32) error {
	return nil
}
