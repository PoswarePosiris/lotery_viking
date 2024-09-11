package models

type RewardView struct {
	RewardID    uint64  `db:"reward_id" json:"reward_id"`
	RewardName  string  `db:"reward_name" json:"reward_name"`
	BigWin      bool    `db:"big_win" json:"big_win"`
	ImageID     *uint64 `db:"image_id" json:"image_id"`
	ImageName   *string `db:"image_name" json:"image_name"`
	ImageFormat *string `db:"image_format" json:"image_format"`
	ImageUrl    *string `db:"image_url" json:"image_url"`
	KioskID     uint64  `db:"kiosk_id" json:"-"`
	ParameterID uint64  `db:"parameter_id" json:"-"`
}
