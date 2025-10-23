package http

import (
	"github.com/labstack/echo/v4"
	"github.com/monforje/user-service/internal/service"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	h.registerUserRoutes(e)
}
