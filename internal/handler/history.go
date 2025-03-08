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
