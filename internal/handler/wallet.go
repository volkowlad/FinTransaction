package handler

import (
	fin "FinTransaction"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// @Summary Create wallet
// @Security ApiKeyAuth
// @Tags wallets
// @Description create wallet
// @ID create-wallet
// @Accept  json
// @Produce  json
// @Param input body FinTransaction.Wallet true "wallet info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /api [post]
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

	log.Printf("Created wallet: %d", id)
}

type getAllWalletsResponse struct {
	Data []fin.Wallet `json:"wallets"`
}

// @Summary Get All Wallets
// @Security ApiKeyAuth
// @Tags wallets
// @Description get all wallets
// @ID get-all-wallets
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllWalletsResponse
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /api [get]
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

// @Summary Get Wallet
// @Security ApiKeyAuth
// @Tags wallets
// @Description get wallet
// @ID get-wallet
// @Accept  json
// @Produce  json
// @Success 200 {object} FinTransaction.Wallet
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /api/:id [get]
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

// @Summary Update Wallet
// @Security ApiKeyAuth
// @Tags wallets
// @Description update wallet
// @ID update-wallet
// @Accept  json
// @Produce  json
// @Success 200 {object} FinTransaction.TransferWallet
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /api/:id [put]
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

// @Summary Delete Wallet
// @Security ApiKeyAuth
// @Tags wallets
// @Description delete wallet
// @ID delete-wallet
// @Accept  json
// @Produce  json
// @Success 200 {object} statusResp
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /api/:id [delete]
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
