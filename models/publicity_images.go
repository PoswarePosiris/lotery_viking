package models

type PublicityImages struct {
	BaseModel
	ParameterId *Parameters `db:"parameter_id" json:"parameter"`
	ImageId     *Images     `db:"image_id" json:"image"`
}
