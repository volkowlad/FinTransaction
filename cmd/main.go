package main

import (
	fin "FinTransaction"
	"FinTransaction/internal/config"
	"FinTransaction/internal/handler"
	"FinTransaction/internal/repository"
	"FinTransaction/internal/service"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
)

// @title Fin Transaction API
// @version 1.0.0
// @description API Server for Fin Application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := config.InitConfig()

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	log := slog.New(slog.NewTextHandler(os.Stdout, opts))

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.User,
		Password: cfg.DB.Password,
		DBName:   cfg.DB.DBName,
		SSLMode:  cfg.DB.SSLMode,
	})
	if err != nil {
		log.Error("failed to initialize database", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(fin.Server)
	go func() {
		if err := srv.Run(cfg.Address, handlers.InitRoutes()); err != nil {
			log.Error("error occured while running http server", err.Error())
		}
	}()

	log.Info("http server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Server is shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Error("server shutdown failed", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Error("failed to close database", err.Error())
	}
}
