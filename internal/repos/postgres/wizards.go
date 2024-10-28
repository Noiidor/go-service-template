package postgres

import (
	"context"

	"github.com/Noiidor/go-service-template/internal/repos"
	"github.com/jmoiron/sqlx"
)

type wizardsRepo struct {
	db *sqlx.DB
}

func NewWizardsRepo(db *sqlx.DB) *wizardsRepo {
	return &wizardsRepo{
		db: db,
	}
}

func (r *wizardsRepo) GetWizardByID(ctx context.Context, id uint32) (*repos.Wizard, error) {
	return nil, nil
}

func (r *wizardsRepo) GetAllWizards(ctx context.Context) ([]*repos.Wizard, error) {
	return nil, nil
}

func (r *wizardsRepo) CreateWizard(ctx context.Context, wizard *repos.Wizard) error {
	return nil
}

func (r *wizardsRepo) UpdateWizard(ctx context.Context, wizard *repos.Wizard) error {
	return nil
}

func (r *wizardsRepo) DeleteWizard(ctx context.Context, id uint32) error {
	return nil
}

func (r *wizardsRepo) AddStatsToWizard(ctx context.Context, stats *repos.WizardStats) error {
	return nil
}

func (r *wizardsRepo) GetWizardStats(ctx context.Context, wizardID uint32) (*repos.WizardStats, error) {
	return nil, nil
}

func (r *wizardsRepo) UpdateStats(ctx context.Context, wizardID uint32, stats *repos.WizardStats) error {
	return nil
}
