package http

import "github.com/labstack/echo/v4"

func (h *Handler) registerUserRoutes(e *echo.Echo) {
	users := e.Group("/api/v1/users")

	users.POST("", h.createUser)
	users.GET("/:telegram_id", h.getUser)
	users.PUT("/:telegram_id", h.updateUser)
	users.DELETE("/:telegram_id", h.deleteUser)
	users.GET("/:telegram_id/exists", h.checkUserExist)
}
