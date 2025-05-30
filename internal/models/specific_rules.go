package models

type SpecificRules struct {
	ID           uint64   `db:"id" json:"id,omitempty"`
	CasinoID     uint64   `db:"casino_id" json:"-"`
	Casino       *Casinos `json:"casino,omitempty"`
	SpecificRule string   `db:"specific_rule" json:"specific_rule"`
}
