package handler

import (
	"database/sql"
	"fmt"
	"lotery_viking/internal/database"
	"lotery_viking/internal/models"
	"lotery_viking/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	BaseHandler
}

const (
	CodeInvalid        = "Code non valide"
	TicketClaimed      = "Ticket déjà réclamé"
	TicketAlreadyAdded = "Ticket déjà ajouté"
	TicketNotFound     = "Ticket non trouvé"
	TicketScanned      = "Ticket déjà scanné"
	TicketCreated      = "Ticket ajouté"
	TicketClaimMsg     = "Ticket réclamé"
	KioskNotFound      = "Kiosk non trouvé"
)

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

	if ticket.EntryScan == nil {
		ticket.SetEntryScanNow()
	}

	// get the id of the kiosk from the context
	macKiosk, exists := c.Request.Context().Value("macKiosk").(string)
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": KioskNotFound})
		return
	}

	// get the id with the func in base BaseHandler
	kiosk, err := t.getKioskView(macKiosk)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusForbidden, gin.H{"error": KioskNotFound})
		return
	}
	clientData := kiosk.ClientData
	ticket.KioskID = kiosk.ID
	// check if the format of the ticket number is correct
	//
	if !ticket.IsValid(kiosk.Secret, kiosk.SecretLength) {
		c.JSON(http.StatusBadRequest, gin.H{"error": CodeInvalid})
		return
	}
	// check if the ticket number is already in the database
	// if ok, return error
	// if not, insert the ticket in the database
	ticketFound, err := t.checkCode(ticket.TicketNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if ticketFound {
		c.JSON(http.StatusConflict, gin.H{"error": TicketAlreadyAdded})
	} else {
		// TODO add client info is need
		db := t.db.GetDB()

		var statement string
		var args []interface{}

		if clientData {
			if !ticket.IsValidClientPhone() {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Numéro de téléphone invalide"})
				return
			}
			statement = "INSERT INTO tickets (kiosk_id, ticket_number, client_phone, entry_scan) VALUES (?, ?,?,?)"
			args = []interface{}{ticket.KioskID, ticket.TicketNumber, ticket.ClientPhone, ticket.EntryScan}
		} else {
			statement = "INSERT INTO tickets (kiosk_id, ticket_number, entry_scan) VALUES (?, ?, ?)"
			args = []interface{}{ticket.KioskID, ticket.TicketNumber, ticket.EntryScan}
		}

		_, err = db.Exec(statement, args...)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": TicketCreated})
	}
}

func (t *TicketHandler) ClaimTicket(c *gin.Context) {
	macAdress, exists := c.Request.Context().Value("macKiosk").(string)
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": KioskNotFound})
	}
	kiosk, err := t.getKioskView(macAdress)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": KioskNotFound})
		return
	}

	code := c.Param("code")

	codeIsValid := utils.DecryptCode(kiosk.Secret, kiosk.SecretLength, code)
	if !codeIsValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": CodeInvalid})
		return
	}

	ticket, err := t.getTicket(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": TicketNotFound})
		return
	}

	if ticket.Claim {
		c.JSON(http.StatusConflict, gin.H{"error": TicketClaimed})
		return
	}

	if ticket.ExitScan != nil {
		c.JSON(http.StatusConflict, gin.H{"error": TicketScanned})
		return
	}

	if ticket.IDReward != nil {
		ticket.Claim = true
	}

	ticket.SetExitScanNow()

	db := t.db.GetDB()
	statement := "UPDATE tickets SET claim = ?, exit_scan = ? WHERE ticket_number = ?"
	_, err = db.Exec(statement, ticket.Claim, ticket.ExitScan, ticket.TicketNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": TicketClaimMsg,
	})
}

func (t *TicketHandler) GetTicket(c *gin.Context) {
	macAdress, exists := c.Request.Context().Value("macKiosk").(string)
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": KioskNotFound})
	}
	kiosk, err := t.getKioskView(macAdress)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": KioskNotFound})
		return
	}

	code := c.Param("code")
	codeIsValid := utils.DecryptCode(kiosk.Secret, kiosk.SecretLength, code)
	if !codeIsValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": CodeInvalid})
		return
	}

	ticket, err := t.getTicket(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": TicketNotFound})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func (t *TicketHandler) getTicket(codeTicket string) (*models.Tickets, error) {
	ticket := &models.Tickets{}
	statement := `
		SELECT t.id, t.kiosk_id, t.id_reward, t.ticket_number,
		       t.client_phone, t.claim, t.entry_scan, t.exit_scan,
		       r.name, r.big_win
		FROM tickets AS t
		LEFT JOIN rewards AS r ON t.id_reward = r.id
		WHERE t.ticket_number = ?`

	db := t.db.GetDB()
	var idReward sql.NullInt64
	var clientPhone sql.NullString
	var exitScan sql.NullTime
	var rewardName sql.NullString
	var bigWin sql.NullBool

	err := db.QueryRow(statement, codeTicket).Scan(
		&ticket.ID,
		&ticket.KioskID,
		&idReward,
		&ticket.TicketNumber,
		&clientPhone,
		&ticket.Claim,
		&ticket.EntryScan,
		&exitScan,
		&rewardName,
		&bigWin,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ticket not found")
		}
		return nil, err
	}

	if !idReward.Valid {
		ticket.IDReward = nil
		ticket.RewardName = nil
		ticket.RewardBigWin = nil
	} else {
		id := uint64(idReward.Int64)
		ticket.IDReward = &id
		ticket.RewardName = &rewardName.String
		ticket.RewardBigWin = &bigWin.Bool
	}

	if !clientPhone.Valid {
		ticket.ClientPhone = nil
	} else {
		ticket.ClientPhone = &clientPhone.String
	}

	if !exitScan.Valid {
		ticket.ExitScan = nil
	} else {
		ticket.ExitScan = &exitScan.Time
	}

	return ticket, nil
}

func (t *TicketHandler) checkCode(codeTicket string) (bool, error) {
	var ticket models.Tickets
	statement := "SELECT id, ticket_number FROM tickets WHERE ticket_number = ?"
	db := t.db.GetDB()
	err := db.QueryRow(statement, codeTicket).Scan(&ticket.ID, &ticket.TicketNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
