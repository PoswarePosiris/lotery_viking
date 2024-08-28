package models

type Rewards struct {
	BaseModel
	Name     string  `db:"name" json:"name"`
	IdImages uint64  `db:"id_images" json:"image,omitempty"`
	Images   *Images `json:"images,omitempty"`
}
