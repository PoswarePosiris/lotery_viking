package models

type Rewards struct {
	BaseModel
	Name     string  `db:"name" json:"name"`
	IdImages *Images `db:"id_images" json:"image,omitempty"`
}
