package service

import (
	"context"

	"github.com/Noiidor/go-service-template/internal/domain"
	"github.com/Noiidor/go-service-template/internal/repos"
)

var _ WizardsService = (*wizardsService)(nil)

type WizardsService interface {
	GetByID(ctx context.Context, id uint32) (*domain.Wizard, error)
	GetAll(ctx context.Context) ([]*domain.Wizard, error)
	Create(ctx context.Context, wizard *domain.Wizard) error
	Update(ctx context.Context, id uint32, wizard *domain.UpdateWizard) (*domain.Wizard, error)
	Delete(ctx context.Context, id uint32) error
	AddStats(ctx context.Context, stats *domain.WizardStats) error
	UpdateStats(ctx context.Context, wizardID uint32, stats *domain.UpdateWizardStats) (*domain.WizardStats, error)
}

type wizardsService struct {
	repo repos.WizardsRepo
}

func NewWizardsService() *wizardsService {
	return &wizardsService{}
}

func (s *wizardsService) GetByID(ctx context.Context, id uint32) (*domain.Wizard, error) {
	wizard, err := s.repo.GetWizardByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return wizard.ToDomain(), nil
}

func (s *wizardsService) GetAll(ctx context.Context) ([]*domain.Wizard, error) {
	wizards, err := s.repo.GetAllWizards(ctx)
	if err != nil {
		return nil, err
	}

	wizardsDomain := make([]*domain.Wizard, len(wizards))
	for i, v := range wizards {
		wizardsDomain[i] = v.ToDomain()
	}

	return wizardsDomain, nil
}

func (s *wizardsService) Create(ctx context.Context, wizard *domain.Wizard) error {
	wiz := new(repos.Wizard)
	wiz.FromDomain(wizard)

	err := s.repo.CreateWizard(ctx, wiz)

	wizard = wiz.ToDomain() // is this a good idea?

	return err
}

func (s *wizardsService) Update(ctx context.Context, id uint32, wizard *domain.UpdateWizard) (*domain.Wizard, error) {
	_, err := s.repo.GetWizardByID(ctx, id)
	if err != nil {
		return nil, err
	}

	wiz := new(repos.UpdateWizard)
	wiz.FromDomain(wizard)

	updated, err := s.repo.UpdateWizard(ctx, id, wiz)
	if err != nil {
		return nil, err
	}

	return updated.ToDomain(), nil
}

func (s *wizardsService) Delete(ctx context.Context, id uint32) error {
	return s.repo.DeleteWizard(ctx, id)
}

func (s *wizardsService) AddStats(ctx context.Context, stats *domain.WizardStats) error {
	_, err := s.repo.GetWizardByID(ctx, stats.WizardID)
	if err != nil {
		return err
	}

	repoStats := new(repos.WizardStats)
	repoStats.FromDomain(stats)

	return s.repo.AddStatsToWizard(ctx, repoStats)
}

func (s *wizardsService) UpdateStats(ctx context.Context, wizardID uint32, stats *domain.UpdateWizardStats) (*domain.WizardStats, error) {
	_, err := s.repo.GetWizardStats(ctx, wizardID)
	if err != nil {
		return nil, err
	}

	stat := new(repos.UpdateWizardStats)
	stat.FromDomain(stats)

	updated, err := s.repo.UpdateStats(ctx, wizardID, stat)
	if err != nil {
		return nil, err
	}

	return updated.ToDomain(), nil
}
