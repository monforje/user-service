package postgres

import (
	"context"
	"time"

	"github.com/monforje/user-service/internal/entity"
)

func (r *UserRepo) CreateUser(ctx context.Context, user *entity.User) (int64, error) {
	query := `
		INSERT INTO users (telegram_id, phone, username, created_at, updated_at)
		VALUES (:telegram_id, :phone, :username, :created_at, :updated_at)
		RETURNING id
	`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	rows, err := r.db.NamedQueryContext(ctx, query, user)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var id int64
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, user *entity.User) error {
	query := `
		UPDATE users
		SET phone = :phone,
		    username = :username,
		    updated_at = :updated_at
		WHERE telegram_id = :telegram_id
	`

	user.UpdatedAt = time.Now()

	_, err := r.db.NamedExecContext(ctx, query, user)
	return err
}

func (r *UserRepo) DeleteUser(ctx context.Context, telegramID int64) error {
	query := `DELETE FROM users WHERE telegram_id = $1`
	_, err := r.db.ExecContext(ctx, query, telegramID)
	return err
}
