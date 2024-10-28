package domain

type Wizard struct {
	ID             uint32 `json:"id"`
	Name           string `json:"name"`
	Specialization string `json:"specialization"`
}

type WizardStats struct {
	WizardID     uint32 `json:"wizardId"`
	Power        int32  `json:"power"`
	Mana         int32  `json:"mana"`
	Intelligence int32  `json:"intelligence"`
	Luck         int32  `json:"luck"`
}
