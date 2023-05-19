package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/serenite11/Links-Reduction-Api/internal/models"
	"math/rand"
	"net/http"
	"time"
)

var symbols = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'_',
}

func (h *Handler) createShortLink(c *gin.Context) {
	var link models.Link
	if err := c.BindJSON(&link); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if _, err := http.Get(link.LongUrl); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Url is not valid!")
		return
	}
	link.ShortUrl = generateShortUrl(link.LongUrl)
	shortUrl, err := h.services.CreateShort(link.LongUrl, link.ShortUrl)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"shortLink": shortUrl,
	})
}

func generateShortUrl(url string) string {
	var shortUrl string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		shortUrl += string(symbols[rand.Intn(52)])
	}
	return shortUrl
}

func (h *Handler) getLink(c *gin.Context) {
	var link models.Link
	if err := c.BindJSON(&link); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	longUrl, err := h.services.GetLongUrl(link.ShortUrl)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"longLink": longUrl,
	})
}
