package postgres

import (
	"context"

	"github.com/Noiidor/go-service-template/internal/repos"
	"github.com/jmoiron/sqlx"
)

var _ repos.WizardsRepo = (*wizardsRepo)(nil)

type wizardsRepo struct {
	db *sqlx.DB
}

func NewWizardsRepo(db *sqlx.DB) *wizardsRepo {
	return &wizardsRepo{
		db: db,
	}
}

func (r *wizardsRepo) GetWizardByID(ctx context.Context, id uint32) (*repos.Wizard, error) {
	wizard := new(repos.Wizard)
	const query = `
		SELECT id, name, specialization
		FROM wizards
		WHERE id = $1;
	`

	err := r.db.GetContext(ctx, wizard, query, id)
	if err != nil {
		return nil, err
	}
	return wizard, nil
}

func (r *wizardsRepo) GetAllWizards(ctx context.Context) ([]*repos.Wizard, error) {
	var wizards []*repos.Wizard
	const query = `
		SELECT id, name, specialization
		FROM wizards;
	`

	err := r.db.SelectContext(ctx, &wizards, query)
	if err != nil {
		return nil, err
	}
	return wizards, nil
}

func (r *wizardsRepo) CreateWizard(ctx context.Context, wizard *repos.Wizard) error {
	const query = `
		INSERT INTO wizards (name, specialization)
		VALUES (:name, :specialization)
		RETURNING id;
	`

	rows, err := r.db.NamedQueryContext(ctx, query, wizard)
	if err != nil {
		return err
	}
	defer rows.Close()

	rows.Scan(wizard.ID)

	return err
}

func (r *wizardsRepo) UpdateWizard(ctx context.Context, wizard *repos.Wizard) error {
	const query = `
		UPDATE wizards 
		SET name = $1, specialization = $2 
		WHERE id = $3;
	`

	_, err := r.db.NamedExecContext(ctx, query, wizard)
	return err
}

func (r *wizardsRepo) DeleteWizard(ctx context.Context, id uint32) error {
	const query = `
		DELETE FROM wizards 
		WHERE id = $1;
	`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *wizardsRepo) AddStatsToWizard(ctx context.Context, stats *repos.WizardStats) error {
	const query = `
		INSERT INTO wizard_stats (wizard_id, power, mana, intelligence, luck) 
		VALUES (:wizard_id, :power, :mana, :intelligence, :luck);
	`

	_, err := r.db.NamedExecContext(ctx, query, stats)
	return err
}

func (r *wizardsRepo) GetWizardStats(ctx context.Context, wizardID uint32) (*repos.WizardStats, error) {
	stats := new(repos.WizardStats)
	const query = `
		SELECT wizard_id, power, mana, intelligence, luck 
		FROM wizard_stats 
		WHERE wizard_id = $1;
	`

	err := r.db.GetContext(ctx, stats, query, wizardID)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

func (r *wizardsRepo) UpdateStats(ctx context.Context, wizardID uint32, stats *repos.WizardStats) error {
	const query = `
		UPDATE wizard_stats
		SET power = :power, mana = :mana, intelligence = :intelligence, luck = :luck
		WHERE wizard_id = :wizard_id;
	`
	_, err := r.db.NamedExecContext(ctx, query, stats)
	return err
}
