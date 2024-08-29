package models

type Rewards struct {
	BaseModel
	Name     string  `db:"name" json:"name"`
	BigWin   bool    `db:"big_win" json:"big_win"`
	IdImages uint64  `db:"id_images" json:"id_images,omitempty"`
	Image    *Images `json:"image,omitempty"`
}
