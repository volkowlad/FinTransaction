package handler

import (
	"FinTransaction/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(RateLimiterMW())
	router.Use(gin.Recovery())

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.signUp)
		auth.POST("/sing-in", h.singIn)
	}

	wallet := router.Group("/api", h.userIdentity)
	{
		wallet.POST("/", h.createWallets)
		wallet.GET("/", h.getAllWallets)
		wallet.GET("/:id", h.getWallet)
		wallet.PUT("/:id", h.updateWallet)
		wallet.DELETE("/:id", h.deleteWallet)
	}

	return router
}
