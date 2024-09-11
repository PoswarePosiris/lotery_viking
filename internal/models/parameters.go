package models

type Status string

const (
	Scan Status = "scan"
	Draw Status = "draw"
)

type Parameters struct {
	BaseModel
	NameLotery    string  `db:"name_lotery" json:"name_lotery"`
	NameCasino    string  `db:"name_casino" json:"name_casino"`
	DateStart     string  `db:"date_start" json:"date_start"`
	DateEnd       string  `db:"date_end" json:"date_end"`
	Status        Status  `db:"status" json:"status"`
	ClientData    bool    `db:"client_data" json:"client_data"`
	HomePageId    uint64  `db:"home_page" json:"home_page_id"`
	HomePage      *Images `json:"home_page,omitempty"`
	ClientPageId  uint64  `db:"client_page" json:"client_page_id"`
	ClientPage    *Images `json:"client_page,omitempty"`
	ResultPageId  uint64  `db:"result_page" json:"result_page_id"`
	ResultPage    *Images `json:"result_page,omitempty"`
	GeneralRules  string  `db:"general_rules" json:"general_rules"`
	SpecificRules string  `db:"specific_rules" json:"specific_rules"`
	Secret        string  `db:"secret" json:"secret"`
	SecretLength  int     `db:"secret_length" json:"secret_length"`
}
