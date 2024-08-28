package handler

import (
	"lotery_viking/internal/database"
	"lotery_viking/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	BaseHandler
}

func NewTicketHandler(db database.Service) *TicketHandler {
	return &TicketHandler{
		BaseHandler: BaseHandler{
			db: db,
		},
	}
}

func (t *TicketHandler) CreateTicket(c *gin.Context) {
	var ticket models.Tickets
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// get the id of the kiosk from the context
	macKiosk := c.Request.Context().Value("macKiosk").(string)

	// get the id with the func in base BaseHandler
	kioskID, err := t.getKioskId(macKiosk)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Kiosk not found"})
		return
	}
	ticket.KioskID.Id = kioskID

	// if err := t.db.CreateTicket(&ticket); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusCreated, ticket)
}

func (t *TicketHandler) checkCode(codeTicket string) (bool, error) {

	statement := "SELECT id FROM tickets WHERE code = ?"
	db := t.db.GetDB()
	err := db.QueryRow(statement, codeTicket).Scan(&id)
	if err != nil {
		return false, err
	}
	return true, nil
}
