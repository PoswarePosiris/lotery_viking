package models

type PublicityImages struct {
	BaseModel
	ParameterId uint64      `db:"parameter_id" json:"parameter_id,omitempty"`
	Parameter   *Parameters `json:"parameter,omitempty"`
	ImageId     uint64      `db:"image_id" json:"image_id,omitempty"`
	Image       *Images     `json:"image,omitempty"`
}
