package models

type PublicityImages struct {
	ID          uint64      `db:"id" json:"id,omitempty"`
	ParameterId uint64      `db:"parameter_id" json:"parameter_id,omitempty"`
	Parameter   *Parameters `json:"parameter,omitempty"`
	KioskId     uint64      `db:"kiosk_id" json:"kiosk_id,omitempty"`
	Kiosk       *Kiosks     `json:"kiosk,omitempty"`
	ImageId     uint64      `db:"image_id" json:"image_id,omitempty"`
	Image       *Images     `json:"image,omitempty"`
}
