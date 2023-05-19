package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/serenite11/Links-Reduction-Api/internal/service"
	"os"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	store := os.Getenv("STORE")
	if store == "POSTGRES" {
		router.POST("/", h.createShortUrl)
		router.GET("/", h.getOriginUrl)
	} else if store == "IN-MEMORY" {
		router.POST("/", h.createShortUrlMemory)
		router.GET("/", h.getOriginUrlMemory)
	}

	return router
}
