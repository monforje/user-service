package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/monforje/user-service/internal/entity"
)

func (h *Handler) getUser(c echo.Context) error {
	telegramIDParam := c.Param("telegram_id")
	telegramID, err := strconv.ParseInt(telegramIDParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid telegram_id"})
	}

	user, err := h.service.User.GetByTelegramID(c.Request().Context(), telegramID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) checkUserExist(c echo.Context) error {
	telegramIDParam := c.Param("telegram_id")

	telegramID, err := strconv.ParseInt(telegramIDParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid telegram_id"})
	}

	exists, err := h.service.User.IsExist(c.Request().Context(), telegramID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, entity.ExistUserResponse{Exists: exists})
}
