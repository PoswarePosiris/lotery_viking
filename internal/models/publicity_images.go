package models

type PublicityImages struct {
	BaseModel
	ParameterId uint64      `db:"parameter_id" json:"omitempty"`
	Parameter   *Parameters `json:"parameter"`
	ImageId     uint64      `db:"image_id" json:"omitempty"`
	Image       *Images     ` json:"image"`
}
