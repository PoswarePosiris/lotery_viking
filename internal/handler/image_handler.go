package handler

import (
	"database/sql"
	"log"
	"lotery_viking/internal/database"
	"lotery_viking/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImagesHandler struct {
	BaseHandler
}

func NewImagesHandler(db database.Service) *ImagesHandler {
	return &ImagesHandler{
		BaseHandler: BaseHandler{
			db: db,
		},
	}
}

func (i *ImagesHandler) GetImages(c *gin.Context) {
	var images []models.Images

	db := i.db.GetDB()
	rows, err := db.Query("SELECT id , name, format , url, created_at, updated_at FROM images")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var image models.Images
		var url sql.NullString
		err = rows.Scan(&image.ID, &image.Name, &image.Format, &url, &image.CreatedAt, &image.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if url.Valid {
			image.Url = &url.String
		} else {
			// Use getPathImage to construct the URL
			fullPath := i.GetPathImage(image.Name, image.Format)
			log.Println(fullPath)
			image.Url = &fullPath
		}
		images = append(images, image)
	}

	err = rows.Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, images)
}

func (i *ImagesHandler) GetImage(c *gin.Context) {
	id := c.Param("id")

	var image models.Images
	var url sql.NullString

	db := i.db.GetDB()
	err := db.QueryRow("SELECT id , name, format , url, created_at, updated_at FROM images WHERE id = ?", id).Scan(&image.ID, &image.Name, &image.Format, &url, &image.CreatedAt, &image.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if url.Valid {
		image.Url = &url.String
	} else {
		// Use getPathImage to construct the URL
		fullPath := i.GetPathImage(image.Name, image.Format)
		log.Println(fullPath)
		image.Url = &fullPath
	}

	c.JSON(http.StatusOK, image)
}
