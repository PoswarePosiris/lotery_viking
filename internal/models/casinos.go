package models

type Casinos struct {
	BaseModel
	Name     string `db:"name" json:"name"`
	Location string `db:"location" json:"location"`
}
