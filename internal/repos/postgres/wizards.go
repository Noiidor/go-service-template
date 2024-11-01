package postgres

import (
	"context"
	"fmt"

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
		return nil, fmt.Errorf("select wizard by ID: %w", err)
	}
	return wizard, nil
}

func (r *wizardsRepo) GetAllWizards(ctx context.Context) ([]*repos.Wizard, error) {
	var wizards []*repos.Wizard
	const query = `
		SELECT id, name, specialization
		FROM wizards
		ORDER BY id;
	`

	err := r.db.SelectContext(ctx, &wizards, query)
	if err != nil {
		return nil, fmt.Errorf("select all wizards: %w", err)
	}
	return wizards, nil
}

func (r *wizardsRepo) CreateWizard(ctx context.Context, wizard *repos.Wizard) error {
	query := `
		INSERT INTO wizards (name, specialization)
		VALUES (:name, :specialization)
		RETURNING id, name, specialization;
	`

	query, args, err := sqlx.Named(query, wizard)
	if err != nil {
		return fmt.Errorf("sqlx named query: %w", err)
	}
	query = r.db.Rebind(query)

	err = r.db.GetContext(ctx, wizard, query, args...)
	if err != nil {
		return fmt.Errorf("insert wizard: %w", err)
	}

	return nil
}

func (r *wizardsRepo) UpdateWizard(ctx context.Context, id uint32, wizard *repos.UpdateWizard) (*repos.Wizard, error) {
	query := `
		UPDATE wizards 
		SET 
			name = COALESCE(:name, name), 
			specialization = COALESCE(:specialization, specialization) 
		WHERE id = ?
		RETURNING id, name, specialization;
	`
	result := new(repos.Wizard)

	query, args, err := sqlx.Named(query, wizard)
	if err != nil {
		return nil, fmt.Errorf("sqlx named query: %w", err)
	}
	query = r.db.Rebind(query)
	args = append(args, id)

	err = r.db.GetContext(ctx, result, query, args...)
	if err != nil {
		return nil, fmt.Errorf("update wizard: %w", err)
	}

	return result, nil
}

func (r *wizardsRepo) DeleteWizard(ctx context.Context, id uint32) error {
	const query = `
		DELETE FROM wizards 
		WHERE id = $1;
	`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete wizard by ID: %w", err)
	}

	return nil
}

func (r *wizardsRepo) AddStatsToWizard(ctx context.Context, stats *repos.WizardStats) error {
	query := `
		INSERT INTO wizard_stats (wizard_id, power, mana, intelligence, luck) 
		VALUES (:wizard_id, :power, :mana, :intelligence, :luck);
	`

	query, args, err := sqlx.Named(query, stats)
	if err != nil {
		return fmt.Errorf("sqlx named query: %w", err)
	}
	query = r.db.Rebind(query)

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("insert wizard stats: %w", err)
	}
	return nil
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
		return nil, fmt.Errorf("select wizard stats: %w", err)
	}
	return stats, nil
}

func (r *wizardsRepo) UpdateStats(ctx context.Context, wizardID uint32, stats *repos.UpdateWizardStats) (*repos.WizardStats, error) {
	query := `
		UPDATE wizard_stats
		SET 
			power = COALESCE(:power, power), 
			mana = COALESCE(:mana, mana), 
			intelligence = COALESCE(:intelligence, intelligence), 
			luck = COALESCE(:luck, luck)
		WHERE wizard_id = ?
		RETURNING wizard_id, power, mana, intelligence, luck;
	`
	result := new(repos.WizardStats)

	query, args, err := sqlx.Named(query, stats)
	if err != nil {
		return nil, fmt.Errorf("sqlx named query: %w", err)
	}
	query = r.db.Rebind(query)
	args = append(args, wizardID)

	err = r.db.GetContext(ctx, result, query, args...)
	if err != nil {
		return nil, fmt.Errorf("update wizard stats: %w", err)
	}
	return result, nil
}
