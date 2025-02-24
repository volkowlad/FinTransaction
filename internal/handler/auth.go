package handler

import (
	fin "FinTransaction"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input fin.User

	if err := c.BindJSON(&input); err != nil {
		NewRespError(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

	slog.Info("user created successfully")
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) singIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		NewRespError(c, http.StatusBadRequest, err.Error())

		return
	}
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

	slog.Info("token generated successfully")
}
