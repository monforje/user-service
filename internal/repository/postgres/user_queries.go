package postgres

import (
	"context"

	"github.com/monforje/user-service/internal/entity"
)

func (r *userRepo) GetByTelegramID(ctx context.Context, telegramID int64) (*entity.User, error) {
	query := `
		SELECT id, telegram_id, phone, username, created_at, updated_at
		FROM users
		WHERE telegram_id = $1
	`

	var user entity.User
	if err := r.db.GetContext(ctx, &user, query, telegramID); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) ExistByTelegramID(ctx context.Context, telegramID int64) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE telegram_id = $1)`

	var exists bool
	if err := r.db.GetContext(ctx, &exists, query, telegramID); err != nil {
		return false, err
	}
	return exists, nil
}
