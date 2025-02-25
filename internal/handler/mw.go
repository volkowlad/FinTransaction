package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"strings"
	"time"
)

const userCtx = "userID"

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		NewRespError(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewRespError(c, http.StatusUnauthorized, "invalid auth header")
	}

	//parse token
	userID, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		NewRespError(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userID)
}

func getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewRespError(c, http.StatusUnauthorized, "user is not found")
		return 0, errors.New("user is not found")
	}

	idInt, ok := id.(int)
	if !ok {
		NewRespError(c, http.StatusInternalServerError, "user id is of invalid type found")
		return 0, errors.New("user id is of invalid type found")
	}

	return idInt, nil
}

func RateLimiterMW() gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(time.Second), 1000)
	return func(c *gin.Context) {

		if limiter.Allow() {
			c.Next()
		} else {
			NewRespError(c, http.StatusTooManyRequests, "too many requests")
		}
	}
}
