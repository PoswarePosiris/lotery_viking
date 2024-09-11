package server

import (
	"lotery_viking/internal/handler"
	"lotery_viking/internal/server/middleware"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

const jsonContentType = "application/json"

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())
	// might remove it in production
	r.Use(gin.Logger())

	// Use the CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "api-key", "accesscontrolalloworigin"}

	r.Use(cors.New(config))

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

	// Group Images
	imageHandler := handler.NewImagesHandler(s.db)
	imageRoutes := r.Group("/images")
	{
		// middleware
		imageRoutes.Use(middleware.CheckAPIKey())

		// routes
		imageRoutes.GET("/", imageHandler.GetImages)
		imageRoutes.GET("/:id", imageHandler.GetImage)
	}

	// Group Kiosks
	kioskHandler := handler.NewKioskHandler(s.db)
	kioskRoutes := r.Group("/kiosks")
	{
		// middleware
		kioskRoutes.Use(middleware.CheckAPIKey())
		// routes
		kioskRoutes.GET("/", kioskHandler.GetKiosk)
		kioskRoutes.Use(middleware.CheckKiosk()).GET("/params", kioskHandler.GetKioskByMac)
	}

	// Add static folder for images
	kioskImagesDir, err := filepath.Abs("kiosk_images")
	if err != nil {
		panic(err)
	}
	r.Static("/kiosk_images", kioskImagesDir)

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
