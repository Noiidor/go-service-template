package repos

import (
	"context"
	"database/sql"

	"github.com/Noiidor/go-service-template/internal/domain"
)

type WizardsRepo interface {
	GetWizardByID(ctx context.Context, id uint32) (*Wizard, error)
	GetAllWizards(ctx context.Context) ([]*Wizard, error)
	CreateWizard(ctx context.Context, wizard *Wizard) error
	UpdateWizard(ctx context.Context, id uint32, wizard *UpdateWizard) (*Wizard, error)
	DeleteWizard(ctx context.Context, id uint32) error

	AddStatsToWizard(ctx context.Context, stats *WizardStats) error
	GetWizardStats(ctx context.Context, wizardID uint32) (*WizardStats, error)
	UpdateStats(ctx context.Context, wizardID uint32, stats *UpdateWizardStats) (*WizardStats, error)
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

type UpdateWizard struct {
	Name           sql.NullString `db:"name"`
	Specialization sql.NullString `db:"specialization"`
}

func (w *UpdateWizard) FromDomain(d *domain.UpdateWizard) {
	if d.Name != nil {
		w.Name.Valid = true
		w.Name.String = *d.Name
	}
	if d.Specialization != nil {
		w.Specialization.Valid = true
		w.Specialization.String = *d.Specialization
	}
}

type WizardStats struct {
	WizardID     uint32 `db:"wizard_id"`
	Power        int32  `db:"power"`
	Mana         int32  `db:"mana"`
	Intelligence int32  `db:"intelligence"`
	Luck         int32  `db:"luck"`
}

type UpdateWizardStats struct {
	Power        sql.NullInt32 `db:"power"`
	Mana         sql.NullInt32 `db:"mana"`
	Intelligence sql.NullInt32 `db:"intelligence"`
	Luck         sql.NullInt32 `db:"luck"`
}

func (s *UpdateWizardStats) FromDomain(d *domain.UpdateWizardStats) {
	if d.Power != nil {
		s.Power.Valid = true
		s.Power.Int32 = *d.Power
	}
	if d.Intelligence != nil {
		s.Intelligence.Valid = true
		s.Intelligence.Int32 = *d.Intelligence
	}
	if d.Mana != nil {
		s.Mana.Valid = true
		s.Mana.Int32 = *d.Mana
	}
	if d.Luck != nil {
		s.Luck.Valid = true
		s.Luck.Int32 = *d.Luck
	}
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
