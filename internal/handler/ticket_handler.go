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
	macKiosk := c.Request.Context().Value("macKiosk").(string)

	// get the id with the func in base BaseHandler
	kiosk, err := t.getKiosk(macKiosk)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusForbidden, gin.H{"error": "Kiosk not found"})
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
		var statement string
		if clientData {
			statement = "INSERT INTO tickets (kiosk_id, ticket_number,client_phone ,entry_scan) VALUES (?, ?,?,  ?)"
		} else {
			statement = "INSERT INTO tickets (kiosk_id, ticket_number, entry_scan) VALUES (?, ?, ?)"
		}
		db := t.db.GetDB()
		_, err = db.Exec(statement, ticket.KioskID, ticket.TicketNumber, ticket.EntryScan)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": TicketCreated})
	}
}

func (t *TicketHandler) ClaimTicket(c *gin.Context) {
	kiosk, err := t.getKiosk(c.Request.Context().Value("macKiosk").(string))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Kiosk not found"})
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
	kiosk, err := t.getKiosk(c.Request.Context().Value("macKiosk").(string))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Kiosk not found"})
		return
	}

	code := c.Param("code")
	fmt.Println(code)
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
		       r.name, r.big_win, i.url
		FROM tickets AS t
		LEFT JOIN rewards AS r ON t.id_reward = r.id
		LEFT JOIN images AS i ON r.id_images = i.id
		WHERE t.ticket_number = ?`

	db := t.db.GetDB()
	var idReward sql.NullInt64
	var clientPhone sql.NullString
	var exitScan sql.NullTime
	var rewardName sql.NullString
	var bigWin sql.NullBool
	var rewardUrl sql.NullString

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
		&rewardUrl,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ticket not found")
		}
		return nil, err
	}

	if !idReward.Valid {
		ticket.RewardName = nil
		ticket.RewardBigWin = nil
		ticket.RewardImage = nil
	} else {
		ticket.RewardName = &rewardName.String
		ticket.RewardBigWin = &bigWin.Bool
		ticket.RewardImage = &rewardUrl.String
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
