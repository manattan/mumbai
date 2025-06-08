package server

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/manattan/mumbai/internal/config"
	"github.com/manattan/mumbai/internal/gateway/repository"
	"github.com/manattan/mumbai/internal/gateway/server/handler"
	"github.com/manattan/mumbai/internal/middleware"
	"github.com/manattan/mumbai/internal/pkg/logger"
	"github.com/manattan/mumbai/internal/usecase"
)

func StartEchoServer(cfg *config.Config) {
	l := logger.New()

	db, err := repository.NewDB(cfg.MySQLDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := repository.NewRepository(db)
	userUseCase := usecase.NewUseCase(repo)
	h := handler.NewHandler(userUseCase, l)

	e := echo.New()

	e.Use(middleware.RequestID)
	e.Use(middleware.Logging(l))

	SetupRoutes(e, h)

	l.Info("Starting Echo server on port %d", cfg.HTTPPort)
	if err := e.Start(fmt.Sprintf(":%d", cfg.HTTPPort)); err != nil {
		l.Error("failed to start server: %v", err)
	}
}