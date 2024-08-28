package handler

import "lotery_viking/internal/database"

type BaseHandler struct {
	db database.Service
}

func (b *BaseHandler) getKioskId(macAddress string) (int, error) {
	var id int
	statement := "SELECT id FROM kiosks WHERE macadress_wifi = ? OR macadress_ethernet = ?"

	db := b.db.GetDB()
	err := db.QueryRow(statement, macAddress, macAddress).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (b *BaseHandler) getKiosk(macAddress string) (int, error) {
	var id int
	statement := "SELECT *  FROM kiosk_view WHERE macadress_wifi = ? OR macadress_ethernet = ?"

	db := b.db.GetDB()
	err := db.QueryRow(statement, macAddress, macAddress).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
