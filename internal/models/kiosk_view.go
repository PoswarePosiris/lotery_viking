package models

import "time"

type KioskView struct {
	ID                  uint64    `db:"id" json:"id"`
	Name                string    `db:"name" json:"name"`
	MacadressWifi       string    `db:"macadress_wifi" json:"macadress_wifi"`
	MacadressEthernet   string    `db:"macadress_ethernet" json:"macadress_ethernet"`
	Location            string    `db:"location" json:"location"`
	NameLotery          string    `db:"name_lotery" json:"name_lotery"`
	NameCasino          string    `db:"name_casino" json:"name_casino"`
	DateStart           string    `db:"date_start" json:"date_start"`
	DateEnd             string    `db:"date_end" json:"date_end"`
	Status              string    `db:"status" json:"status"`
	ClientData          bool      `db:"client_data" json:"client_data"`
	Publicity           *string   `db:"publicity" json:"publicity"`
	HomePageId          *uint64   `db:"home_page" json:"-"`
	HomePage            *Images   `json:"home_page"`
	ClientPageId        *uint64   `db:"client_page" json:"-"`
	ClientPage          *Images   `json:"client_page"`
	ResultPageId        *uint64   `db:"result_page" json:"-"`
	ResultPage          *Images   `json:"result_page"`
	GeneralRules        string    `db:"general_rules" json:"general_rules"`
	SpecificRules       *string   `db:"specific_rules" json:"specific_rules"`
	Secret              string    `db:"secret" json:"secret"`
	SecretLength        int       `db:"secret_length" json:"secret_length"`
	UpdatedAt           time.Time `db:"updated_at" json:"updated_at"`
	UpdatedAtParameters time.Time `db:"updated_at_parameters" json:"updated_at_parameters"`
}
