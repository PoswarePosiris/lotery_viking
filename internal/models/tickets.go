package models

import (
	"lotery_viking/internal/utils"
	"time"
)

type Tickets struct {
	BaseModel
	KioskID      uint64     `db:"kiosk_id" json:"kiosk_id"`
	Kiosk        *Kiosks    `json:"kiosk,omitempty"`
	IDReward     *uint64    `db:"id_reward" json:"id_reward"`
	TicketNumber string     `db:"ticket_number" json:"ticket_number"`
	ClientPhone  *string    `db:"client_phone" json:"client_phone"`
	Claim        bool       `db:"claim" json:"claim"`
	EntryScan    *time.Time `db:"entry_scan" json:"entry_scan"`
	ExitScan     *time.Time `db:"exit_scan" json:"exit_scan"`
	RewardName   *string    `json:"reward_name"`
	RewardBigWin *bool      `json:"big_win"`
	RewardImage  *string    `json:"reward_image"`
}

func (t *Tickets) SetEntryScanNow() {
	now := time.Now()
	t.EntryScan = &now
}

func (t *Tickets) SetExitScanNow() {
	now := time.Now()
	t.ExitScan = &now
}

func (t *Tickets) isClaimed() bool {
	return t.Claim
}

func (t *Tickets) IsValid(secret string, secretLength int) bool {
	return utils.DecryptCode(secret, secretLength, t.TicketNumber)
}

func (t *Tickets) IsValidClientPhone() bool {
	return utils.ValidatePhoneNumber(*t.ClientPhone)
}
