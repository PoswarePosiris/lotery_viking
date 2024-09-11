package handler

import (
	"database/sql"
	"log"
	"lotery_viking/internal/database"
	"lotery_viking/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KioskHandler struct {
	BaseHandler
}

func NewKioskHandler(db database.Service) *KioskHandler {
	return &KioskHandler{
		BaseHandler: BaseHandler{
			db: db,
		},
	}
}

func (k *KioskHandler) GetKiosk(c *gin.Context) {
	var kiosks []models.Kiosks

	db := k.db.GetDB()

	rows, err := db.Query("SELECT id , name , macadress_wifi, macadress_ethernet , location , id_parameters , created_at, updated_at FROM kiosks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var kiosk models.Kiosks
		err := rows.Scan(&kiosk.ID, &kiosk.Name, &kiosk.MacadressWifi, &kiosk.MacadressEthernet, &kiosk.Location, &kiosk.IdParameters, &kiosk.CreatedAt, &kiosk.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		kiosks = append(kiosks, kiosk)
	}

	err = rows.Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, kiosks)
}

func (k *KioskHandler) GetKioskByMac(c *gin.Context) {
	mac, ok := k.getMac(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid MAC address"})
		return
	}

	// Fetch the Kiosk by MAC address
	kiosk, err := k.getKioskView(mac)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Kiosk not found"})
		return
	}

	// Collect all image IDs from KioskView and fetch images in one go
	imageIds := collectImageIds(kiosk)
	// Fetch additional image IDs from publicity and rewards
	publicityIds, err := k.getPublicityImageIds(kiosk.ID, kiosk.ParametersID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch additional image IDs"})
		return
	}

	// Combine and remove duplicate IDs
	uniqueImageIds := removeDuplicateIds(append(imageIds, publicityIds...))
	// Fetch all images using a single call
	imagesMap, err := k.getImagesFromList(uniqueImageIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch images"})
		return
	}

	// Assign the fetched images back to the Kiosk object
	assignImagesToKiosk(kiosk, imagesMap)

	// Assign publicity images
	assignPublicityImages(kiosk, imagesMap, publicityIds)

	// Return the populated kiosk object
	c.JSON(http.StatusOK, kiosk)
}

// Function to get all rewards images from the kiosk
func (k *KioskHandler) GetKioskRewards(c *gin.Context) {
	mac, ok := k.getMac(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid MAC address"})
		return
	}

	kiosk, err := k.getKiosk(mac)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Kiosk not found"})
		return
	}

	db := k.db.GetDB()

	statement := "SELECT reward_id, reward_name, big_win, image_id, image_name, image_format, image_url FROM reward_view WHERE kiosk_id = ? OR parameter_id = ?"
	// Fetch IDs from publicity
	rows, err := db.Query(statement, kiosk.ID, kiosk.IdParameters)
	if err != nil {
		log.Println("Error fetching rewards ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch rewards"})
		return
	}

	defer rows.Close()

	var rewards []models.RewardView
	for rows.Next() {
		var reward models.RewardView
		var imageIDIsNull sql.NullInt64
		var imageName, imageFormat, imageUrl sql.NullString
		err := rows.Scan(&reward.RewardID, &reward.RewardName, &reward.BigWin, &imageIDIsNull, &imageName, &imageFormat, &imageUrl)
		if err != nil {
			log.Println("Error scanning rewards ", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch rewards"})
			return
		}
		if imageIDIsNull.Valid {
			imageID := uint64(imageIDIsNull.Int64)
			reward.ImageID = &imageID
		} else {
			continue
		}

		if imageName.Valid {
			reward.ImageName = &imageName.String
		}
		if imageFormat.Valid {
			reward.ImageFormat = &imageFormat.String
		}

		if imageUrl.Valid {
			reward.ImageUrl = &imageUrl.String
		}
		if !imageUrl.Valid && reward.ImageName != nil && reward.ImageFormat != nil {
			fullPath := k.GetPathImage(*reward.ImageName, *reward.ImageFormat)
			reward.ImageUrl = &fullPath
		}
		rewards = append(rewards, reward)
	}

	c.JSON(http.StatusOK, rewards)
}

// Helper function to remove duplicate IDs
func removeDuplicateIds(ids []uint64) []uint64 {
	uniqueIds := make(map[uint64]struct{})
	for _, id := range ids {
		uniqueIds[id] = struct{}{}
	}

	var result []uint64
	for id := range uniqueIds {
		result = append(result, id)
	}

	return result
}

// Helper function to collect image IDs from KioskView
func collectImageIds(kiosk *models.KioskView) []uint64 {
	var ids []uint64
	if kiosk.HomePageId != nil {
		ids = append(ids, *kiosk.HomePageId)
	}
	if kiosk.ClientPageId != nil {
		ids = append(ids, *kiosk.ClientPageId)
	}
	if kiosk.ResultPageId != nil {
		ids = append(ids, *kiosk.ResultPageId)
	}
	// Add more IDs if necessary
	return ids
}

// Helper function to assign fetched images back to KioskView
func assignImagesToKiosk(kiosk *models.KioskView, imagesMap map[uint64]models.Images) {
	if kiosk.HomePageId != nil {
		if image, ok := imagesMap[*kiosk.HomePageId]; ok {
			kiosk.HomePage = &image
		}
	}
	if kiosk.ClientPageId != nil {
		if image, ok := imagesMap[*kiosk.ClientPageId]; ok {
			kiosk.ClientPage = &image
		}
	}
	if kiosk.ResultPageId != nil {
		if image, ok := imagesMap[*kiosk.ResultPageId]; ok {
			kiosk.ResultPage = &image
		}
	}
}

// Helper function to assign publicity images
func assignPublicityImages(kiosk *models.KioskView, imagesMap map[uint64]models.Images, publicityIds []uint64) {
	var publicityImages []models.Images
	for _, id := range publicityIds {
		if image, ok := imagesMap[id]; ok {
			publicityImages = append(publicityImages, image)
		}
	}
	kiosk.Publicity = &publicityImages
}

// getPublicityImageIds fetches IDs from the publicity
func (k *KioskHandler) getPublicityImageIds(kioksId uint64, parametersId uint64) ([]uint64, error) {
	var ids []uint64

	db := k.db.GetDB()

	statement := "SELECT image_id FROM publicity_images WHERE kiosk_id = ? OR parameter_id = ?"
	// Fetch IDs from publicity
	rows, err := db.Query(statement, kioksId, parametersId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id uint64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}
