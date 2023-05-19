package handlers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Error struct {
	message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Error(message)
	c.AbortWithStatusJSON(statusCode, Error{message})
}
