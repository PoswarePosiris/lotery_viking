package handler

import (
	"database/sql"
	"fmt"
	"lotery_viking/internal/database"
	"lotery_viking/internal/models"
	"os"
	"strings"
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

func (b *BaseHandler) getKiosk(macAddress string) (*models.KioskView, error) {
	// var kiosk models.KioskView
	kiosk := &models.KioskView{}
	statement := "SELECT id, name ,macadress_wifi, macadress_ethernet , location , name_lotery , name_casino , date_start  , date_end  , status , client_data , home_page, client_page , result_page ,general_rules,specific_rules, secret  , secret_length , updated_at , updated_at_parameters FROM kiosk_view WHERE macadress_wifi = ? OR macadress_ethernet = ?"

	db := b.db.GetDB()
	var homePageIdNull, clientPageIdNull, resultPageIdNull sql.NullInt64
	var specificRulesNull sql.NullString
	err := db.QueryRow(statement, macAddress, macAddress).Scan(&kiosk.ID, &kiosk.Name, &kiosk.MacadressWifi, &kiosk.MacadressEthernet, &kiosk.Location, &kiosk.NameLotery, &kiosk.NameCasino, &kiosk.DateStart, &kiosk.DateEnd, &kiosk.Status, &kiosk.ClientData, &homePageIdNull, &clientPageIdNull, &resultPageIdNull, &kiosk.GeneralRules, &specificRulesNull, &kiosk.Secret, &kiosk.SecretLength, &kiosk.UpdatedAt, &kiosk.UpdatedAtParameters)
	if err != nil {
		return nil, err
	}
	if homePageIdNull.Valid {
		homePageId := uint64(homePageIdNull.Int64)
		kiosk.HomePageId = &homePageId
	}
	if clientPageIdNull.Valid {
		clientPageId := uint64(clientPageIdNull.Int64)
		kiosk.ClientPageId = &clientPageId
	}
	if resultPageIdNull.Valid {
		resultPageId := uint64(resultPageIdNull.Int64)
		kiosk.ResultPageId = &resultPageId
	}
	if specificRulesNull.Valid {
		kiosk.SpecificRules = &specificRulesNull.String
	}

	return kiosk, nil
}

func (b *BaseHandler) GetPathImage(name string, format string) string {
	haveSSL := os.Getenv("HTTP")
	basePath := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if haveSSL == "" {
		haveSSL = "http"
	}

	if basePath == "" {
		basePath = "localhost"
	}
	if port == "" {
		port = "8080"
	}

	return fmt.Sprintf("%s://%s:%s/kiosk_images/%s.%s", haveSSL, basePath, port, name, format)
}

func (b *BaseHandler) getImagesFromList(list []uint64) ([]models.Images, error) {
	var images []models.Images

	placeholders := make([]string, len(list))
	for i := range list {
		placeholders[i] = "?"
	}

	query := fmt.Sprintf("SELECT id, name, url FROM images WHERE id IN (%s)", strings.Join(placeholders, ", "))

	// Prepare the query
	db := b.db.GetDB()
	rows, err := db.Query(query, convertUint64SliceToInterfaceSlice(list)...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var image models.Images
		var url sql.NullString
		err = rows.Scan(&image.ID, &image.Name, &url, &image.Format)
		if err != nil {
			return nil, err
		}
		if url.Valid {
			image.Url = &url.String
		} else {
			// Use getPathImage to construct the URL
			fullPath := b.GetPathImage(image.Name, image.Format)
			image.Url = &fullPath
		}

		images = append(images, image)
	}

	// Check for errors encountered during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return images, nil
}

// Helper function to convert []uint64 to []interface{} for query
func convertUint64SliceToInterfaceSlice(slice []uint64) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}
