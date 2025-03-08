package handler

import (
	"FinTransaction/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "FinTransaction/docs"
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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.singIn)
	}

	wallet := router.Group("/api", h.userIdentity)
	{
		wallet.POST("/", h.createWallets)
		wallet.GET("/", h.getAllWallets)
		wallet.GET("/:id", h.getWallet)
		wallet.PUT("/:id", h.updateWallet)
		wallet.DELETE("/:id", h.deleteWallet)
		wallet.GET("/history/:id", h.historyWallet)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
