package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/serenite11/Links-Reduction-Api/internal/models"
	"net/http"
)

func (h *Handler) createShortUrl(c *gin.Context) {
	var link models.Link
	if err := c.BindJSON(&link); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := validateCreateShortUrl(link.LongUrl); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	shortUrl, err := h.services.LinksShortener.CreateShortUrl(link.LongUrl)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"shortLink": shortUrl,
	})
}

func (h *Handler) getOriginUrl(c *gin.Context) {
	var link models.Link
	if err := c.BindJSON(&link); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	longUrl, err := h.services.LinksShortener.GetLongUrl(link.ShortUrl)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"longLink": longUrl,
	})
}
