package server

import (
	"lotery_viking/internal/handler"
	"lotery_viking/internal/server/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

const jsonContentType = "application/json"

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.Use()
	// public route
	r.GET("/", s.HelloWorldHandler)

	// protected route
	r.GET("/test", middleware.CheckAPIKey(), s.HelloWorldHandler)

	// Health check
	r.GET("/health", s.healthHandler)

	// Group Ticket
	ticketHandler := handler.NewTicketHandler(s.db)
	ticketRoutes := r.Group("/tickets")
	{
		// middleware
		ticketRoutes.Use(middleware.CheckAPIKey())
		ticketRoutes.Use(middleware.CheckKiosk())

		// routes
		ticketRoutes.POST("/", ticketHandler.CreateTicket)
		ticketRoutes.GET("/:code", ticketHandler.GetTicket)
		ticketRoutes.GET("/claim/:code", ticketHandler.ClaimTicket)
	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
