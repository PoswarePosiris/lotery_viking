package models

import "time"

type Tickets struct {
	BaseModel
	KioskID      *Kiosks    `db:"kiosk_id" json:"kiosk_id"`
	IDReward     *Rewards   `db:"id_reward" json:"id_reward,omitempty"`
	TicketNumber string     `db:"ticket_number" json:"ticket_number"`
	Claim        bool       `db:"claim" json:"claim"`
	EntryScan    time.Time  `db:"entry_scan" json:"entry_scan"`
	ExitScan     *time.Time `db:"exit_scan" json:"exit_scan,omitempty"`
}

func (t Tickets) isClaim() {

}
