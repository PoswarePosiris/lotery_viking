package models

type SpecificRules struct {
	ID           uint64  `db:"id" json:"id,omitempty"`
	KioskId      uint64  `db:"kiosk_id" json:"kiosk_id,omitempty"`
	Kiosk        *Kiosks `json:"kiosk,omitempty"`
	SpecificRule string  `db:"specific_rule" json:"specific_rule"`
}
