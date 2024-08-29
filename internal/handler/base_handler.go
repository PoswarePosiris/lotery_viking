package handler

import (
	"lotery_viking/internal/database"
	"lotery_viking/internal/models"
)

type BaseHandler struct {
	db database.Service
}

func (b *BaseHandler) getKioskId(macAddress string) (uint64, error) {
	var id uint64
	statement := "SELECT id FROM kiosks WHERE macadress_wifi = ? OR macadress_ethernet = ?"

	db := b.db.GetDB()
	err := db.QueryRow(statement, macAddress, macAddress).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (b *BaseHandler) getKiosk(macAddress string) (models.KioskView, error) {
	var kiosk models.KioskView
	statement := "SELECT id, name ,macadress_wifi, macadress_ethernet , location , name_lotery , name_casino , date_start  , date_end  , status , client_data , publicity, home_page, scan_page , result_page ,general_rules,specific_rules, secret  , secret_length , updated_at , updated_at_parameters FROM kiosk_view WHERE macadress_wifi = ? OR macadress_ethernet = ?"

	db := b.db.GetDB()
	err := db.QueryRow(statement, macAddress, macAddress).Scan(&kiosk.ID, &kiosk.Name, &kiosk.MacadressWifi, &kiosk.MacadressEthernet, &kiosk.Location, &kiosk.NameLotery, &kiosk.NameCasino, &kiosk.DateStart, &kiosk.DateEnd, &kiosk.Status, &kiosk.ClientData, &kiosk.Publicity, &kiosk.HomePage, &kiosk.ScanPage, &kiosk.ResultPage, &kiosk.GeneralRules, &kiosk.SpecificRules, &kiosk.Secret, &kiosk.SecretLength, &kiosk.UpdatedAt, &kiosk.UpdatedAtParameters)
	if err != nil {
		return kiosk, err
	}
	return kiosk, nil
}
