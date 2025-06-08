package server

import (
	"github.com/labstack/echo/v4"

	"github.com/manattan/mumbai/internal/gateway/server/handler"
)

func SetupRoutes(e *echo.Echo, h handler.Handler) {
	api := e.Group("/api/v1")

	users := api.Group("/users")
	users.POST("", h.CreateUser)
	users.GET("/:id", h.GetUser)
	users.GET("", h.ListUsers)
}
