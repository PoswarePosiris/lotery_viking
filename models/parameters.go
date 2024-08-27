package models

type Parameters struct {
	BaseModel
	NameLotery    string `db:"name_lotery" json:"name_lotery"`
	NameCasino    string `db:"name_casino" json:"name_casino"`
	DateStart     string `db:"date_start" json:"date_start"`
	DateEnd       string `db:"date_end" json:"date_end"`
	Status        string `db:"status" json:"status"`
	HomePage      uint   `db:"home_page" json:"home_page"`
	ScanPage      uint   `db:"scan_page" json:"scan_page"`
	ResultPage    uint   `db:"result_page" json:"result_page"`
	GeneralRules  string `db:"general_rules" json:"general_rules"`
	SpecificRules string `db:"specific_rules" json:"specific_rules"`
	Secret        string `db:"secret" json:"secret"`
	SecretLength  int    `db:"secret_length" json:"secret_length"`
}
