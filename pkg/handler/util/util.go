package util

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/* Структура сообщения об ошибке*/
type ResponseMessage struct {
	Message string `json:"message" binding:"required"`
}

/* Генерация сообщений об ошибке */
func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, ResponseMessage{Message: message})
}
