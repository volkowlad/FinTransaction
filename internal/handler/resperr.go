package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

type respError struct {
	Massage string `json:"massage"`
}

type statusResp struct {
	Status string `json:"status"`
}

func NewRespError(c *gin.Context, statusCode int, message string) {
	slog.Error(message)
	c.AbortWithStatusJSON(statusCode, respError{message})
}
