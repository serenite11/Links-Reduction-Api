package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/serenite11/Links-Reduction-Api/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/", h.createShortLink)
	router.GET("/", h.getLink)

	return router
}
