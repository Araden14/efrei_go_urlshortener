package api

import (
	"github.com/gin-gonic/gin"
	"github.com/axellelanca/urlshortener/internal/services"
)

type Handler struct {
	Service *services.LinkService
}

func NewHandler(service *services.LinkService) *Handler {
	return &Handler{Service: service}
}

// RegisterRoutes définit les routes (bouchon)
func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}