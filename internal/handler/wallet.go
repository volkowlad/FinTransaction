package handler

import (
	fin "FinTransaction"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) createWallets(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var input fin.Wallet
	if err := c.BindJSON(&input); err != nil {
		NewRespError(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.CreateWallet(userID, input)
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

	slog.Info("Created wallet: ", id)
}

type getAllWalletsResponse struct {
	Data []fin.Wallet `json:"wallets"`
}

func (h *Handler) getAllWallets(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	wallets, err := h.services.GetAllWallets(userID)
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllWalletsResponse{
		Data: wallets,
	})

	fmt.Printf("take all wallets: %d", userID)
}

func (h *Handler) getWallet(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	wallet, err := h.services.GetIDWallet(userID, id)
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, wallet)

	fmt.Printf("take wallet: %d", id)
}

func (h *Handler) updateWallet(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input fin.TransferWallet
	if err := c.BindJSON(&input); err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	balance, err := h.services.Transfer(userID, id, input)
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"new balance": balance,
	})

	fmt.Printf("new balance: %d", id)
}

func (h *Handler) deleteWallet(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.DeleteIDWallet(userID, id)
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResp{
		Status: "ok",
	})
}
