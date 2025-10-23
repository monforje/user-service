package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/monforje/user-service/internal/entity"
)

func (h *Handler) createUser(c echo.Context) error {
	var req entity.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	user := &entity.User{
		TelegramID: req.TelegramID,
		Phone:      req.Phone,
		Username:   req.Username,
	}

	id, err := h.service.User.Create(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, entity.CreateUserResponse{ID: id})
}

func (h *Handler) updateUser(c echo.Context) error {
	telegramIDParam := c.Param("telegram_id")
	telegramID, err := strconv.ParseInt(telegramIDParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid telegram_id"})
	}

	var user entity.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid body"})
	}

	user.TelegramID = telegramID
	if err := h.service.User.Update(c.Request().Context(), &user); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "updated"})
}

func (h *Handler) deleteUser(c echo.Context) error {
	telegramIDParam := c.Param("telegram_id")
	telegramID, err := strconv.ParseInt(telegramIDParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid telegram_id"})
	}

	if err := h.service.User.Delete(c.Request().Context(), telegramID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
