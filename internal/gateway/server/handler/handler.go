package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/manattan/mumbai/internal/pkg/logger"
	"github.com/manattan/mumbai/internal/response"
	"github.com/manattan/mumbai/internal/usecase"
)

type Handler interface {
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
	ListUsers(c echo.Context) error
}

type handler struct {
	useCase usecase.UseCase
	logger  logger.Logger
}

func NewHandler(useCase usecase.UseCase, logger logger.Logger) Handler {
	return &handler{
		useCase: useCase,
		logger:  logger,
	}
}

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *handler) CreateUser(c echo.Context) error {
	var req CreateUserRequest
	if err := c.Bind(&req); err != nil {
		h.logger.Error("failed to bind request: %v", err)
		return c.JSON(http.StatusBadRequest, response.NewAppError("INVALID_REQUEST", "Invalid request format", nil))
	}

	user, err := h.useCase.CreateUser(c.Request().Context(), req.Name, req.Email)
	if err != nil {
		h.logger.Error("failed to create user: %v", err)
		return c.JSON(http.StatusInternalServerError, response.NewAppError("CREATE_USER_FAILED", "Failed to create user", nil))
	}

	return c.JSON(http.StatusCreated, UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}

func (h *handler) GetUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.logger.Error("invalid user id: %s", idStr)
		return c.JSON(http.StatusBadRequest, response.NewAppError("INVALID_USER_ID", "Invalid user ID", nil))
	}

	user, err := h.useCase.GetUser(c.Request().Context(), uint(id))
	if err != nil {
		h.logger.Error("failed to get user: %v", err)
		return c.JSON(http.StatusNotFound, response.NewAppError("USER_NOT_FOUND", "User not found", nil))
	}

	return c.JSON(http.StatusOK, UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}

func (h *handler) ListUsers(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")

	limit := 10
	offset := 0

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
			offset = o
		}
	}

	users, err := h.useCase.ListUsers(c.Request().Context(), limit, offset)
	if err != nil {
		h.logger.Error("failed to list users: %v", err)
		return c.JSON(http.StatusInternalServerError, response.NewAppError("LIST_USERS_FAILED", "Failed to list users", nil))
	}

	userResponses := make([]UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
	}

	return c.JSON(http.StatusOK, userResponses)
}
