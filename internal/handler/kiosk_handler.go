package handler

import (
	"fmt"
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
	r := c.Request
	if r == nil {
		fmt.Println("Request is nil")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mac address"})
		return
	}
	// check if the mac address is in the context
	ctx := r.Context()
	if ctx == nil {
		fmt.Println("Context is nil")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mac address"})
		return
	}

	mac := ctx.Value("macKiosk").(string)
	if mac == "" {
		fmt.Println("mac is nil")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mac address"})
		return
	}
	fmt.Println(mac)
	kiosk, err := k.getKiosk(mac)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Kiosk not found"})
		return
	}

	c.JSON(http.StatusOK, kiosk)
}
