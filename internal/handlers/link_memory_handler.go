package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type link struct {
	longUrl  string `json:"long_url"`
	shortUrl string `json:"short_url"`
}

func (h *Handler) createShortUrlMemory(c *gin.Context) {
	var link link
	if err := c.BindJSON(&link); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	shortUrl := h.services.CreateShortUrlMemory(link.longUrl, GenerateShortUrl(link.shortUrl))
	c.JSON(http.StatusOK, map[string]interface{}{
		"shortLink": shortUrl,
	})
}

func (h *Handler) getOriginUrlMemory(c *gin.Context) {
	var link link
	if err := c.BindJSON(&link); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	longUrl := h.services.GetLongUrlMemory(GenerateShortUrl(link.longUrl))
	c.JSON(http.StatusOK, map[string]interface{}{
		"longUrl": longUrl,
	})
}
