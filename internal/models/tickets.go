package models

import "time"

type Tickets struct {
	BaseModel
	KioskID      uint64     `db:"kiosk_id" json:"kiosk_id"`
	Kiosk        *Kiosks    `json:"kiosk,omitempty"`
	IDReward     uint64     `db:"id_reward" json:"id_reward,omitempty"`
	Reward       *Rewards   `json:"reward,omitempty"`
	TicketNumber string     `db:"ticket_number" json:"ticket_number"`
	Claim        bool       `db:"claim" json:"claim"`
	EntryScan    time.Time  `db:"entry_scan" json:"entry_scan"`
	ExitScan     *time.Time `db:"exit_scan" json:"exit_scan,omitempty"`
}

func (t Tickets) SetEntryScanNow() {
	t.EntryScan = time.Now()
}

func (t Tickets) SetExitScanNow() {
	now := time.Now()
	t.ExitScan = &now
}

func (t Tickets) isClaim() {

}
