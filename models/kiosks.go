package models

type Kiosks struct {
	BaseModel
	Name              string      `db:"name" json:"name"`
	MacadressWifi     string      `db:"macadress_wifi" json:"macadress_wifi"`
	MacadressEthernet string      `db:"macadress_ethernet" json:"macadress_ethernet"`
	Location          string      `db:"location" json:"location"`
	IdParameters      *Parameters `db:"id_parameters" json:"parameters,omitempty"`
}
