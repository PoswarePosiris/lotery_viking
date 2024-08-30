package models

import "time"

type KioskView struct {
	ID                  uint64    `db:"id"`
	Name                string    `db:"name"`
	MacadressWifi       string    `db:"macadress_wifi"`
	MacadressEthernet   string    `db:"macadress_ethernet"`
	Location            string    `db:"location"`
	NameLotery          string    `db:"name_lotery"`
	NameCasino          string    `db:"name_casino"`
	DateStart           string    `db:"date_start"`
	DateEnd             string    `db:"date_end"`
	Status              string    `db:"status"`
	ClientData          bool      `db:"client_data"`
	Publicity           *string   `db:"publicity"`
	HomePage            *string   `db:"home_page"`
	ScanPage            *string   `db:"scan_page"`
	ResultPage          *string   `db:"result_page"`
	GeneralRules        string    `db:"general_rules"`
	SpecificRules       *string   `db:"specific_rules"`
	Secret              string    `db:"secret"`
	SecretLength        int       `db:"secret_length"`
	UpdatedAt           time.Time `db:"updated_at"`
	UpdatedAtParameters time.Time `db:"updated_at_parameters"`
}
