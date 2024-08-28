package models

type Kiosks struct {
	BaseModel
	Name              string      `db:"name" json:"name"`
	MacadressWifi     string      `db:"macadress_wifi" json:"macadress_wifi"`
	MacadressEthernet string      `db:"macadress_ethernet" json:"macadress_ethernet"`
	Location          string      `db:"location" json:"location"`
	IdParameters      uint64      `db:"id_parameters" json:"id_parameters,omitempty"` // Correctly maps to the foreign key ID
	Parameters        *Parameters `json:"parameters,omitempty"`
}
