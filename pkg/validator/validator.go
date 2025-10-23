package validator

import (
	"errors"
	"regexp"
	"strings"
)

var (
	phoneRegex    = regexp.MustCompile(`^(\+7|8)\d{10}$`)
	usernameRegex = regexp.MustCompile(`^@[a-zA-Z0-9_]{5,32}$`)
)

func ValidateTelegramID(telegramID int64) error {
	if telegramID <= 0 {
		return errors.New("invalid telegram_id")
	}
	return nil
}

func ValidatePhone(phone *string) error {
	if phone == nil {
		return nil
	}

	cleaned := strings.ReplaceAll(*phone, " ", "")
	cleaned = strings.ReplaceAll(cleaned, "-", "")
	cleaned = strings.ReplaceAll(cleaned, "(", "")
	cleaned = strings.ReplaceAll(cleaned, ")", "")

	if !phoneRegex.MatchString(cleaned) {
		return errors.New("invalid phone number")
	}

	return nil
}

func ValidateUsername(username *string) error {
	if username == nil {
		return nil
	}

	if !usernameRegex.MatchString(*username) {
		return errors.New("invalid username")
	}

	return nil
}
