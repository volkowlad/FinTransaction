package handler

import (
	fin "FinTransaction"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type historyResponse struct {
	Data []fin.History `json:"history"`
}

// @Summary Get Wallet History
// @Security ApiKeyAuth
// @Tags history
// @Description wallet history
// @ID wallet-history
// @Accept  json
// @Produce  json
// @Success 200 {object} FinTransaction.History
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /api/history/:id [get]
func (h *Handler) historyWallet(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	history, err := h.services.HistoryWallet(userID)
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, historyResponse{
		Data: history,
	})

	slog.Info("take history wallet", userID)
}
