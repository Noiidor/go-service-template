package repos

import (
	"context"

	"github.com/Noiidor/go-service-template/internal/domain"
)

type WizardsRepo interface {
	GetWizardByID(ctx context.Context, id uint32) (*Wizard, error)
	GetAllWizards(ctx context.Context) ([]*Wizard, error)
	CreateWizard(ctx context.Context, wizard *Wizard) error
	UpdateWizard(ctx context.Context, wizard *Wizard) error
	DeleteWizard(ctx context.Context, id uint32) error

	AddStatsToWizard(ctx context.Context, stats *WizardStats) error
	GetWizardStats(ctx context.Context, wizardID uint32) (*WizardStats, error)
	UpdateStats(ctx context.Context, wizardID uint32, stats *WizardStats) error
}

type Wizard struct {
	ID             uint32 `db:"id"`
	Name           string `db:"name"`
	Specialization string `db:"specialization"`
}

func (w *Wizard) ToDomain() *domain.Wizard {
	return &domain.Wizard{
		ID:             w.ID,
		Name:           w.Name,
		Specialization: w.Specialization,
	}
}

func (w *Wizard) FromDomain(d *domain.Wizard) {
	*w = Wizard{ // hmmm...
		ID:             d.ID,
		Name:           d.Name,
		Specialization: d.Specialization,
	}
}

type WizardStats struct {
	WizardID     uint32 `db:"wizard_id"`
	Power        int32  `db:"power"`
	Mana         int32  `db:"mana"`
	Intelligence int32  `db:"intelligence"`
	Luck         int32  `db:"luck"`
}

func (ws *WizardStats) ToDomain() *domain.WizardStats {
	return &domain.WizardStats{
		WizardID:     ws.WizardID,
		Power:        ws.Power,
		Mana:         ws.Mana,
		Intelligence: ws.Intelligence,
		Luck:         ws.Luck,
	}
}

func (ws *WizardStats) FromDomain(d *domain.WizardStats) {
	*ws = WizardStats{
		WizardID:     d.WizardID,
		Power:        d.Power,
		Mana:         d.Mana,
		Intelligence: d.Intelligence,
		Luck:         d.Luck,
	}
}
