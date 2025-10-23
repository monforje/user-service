package entity

import "time"

type User struct {
	ID         int64     `db:"id"`
	TelegramID int64     `db:"telegram_id"`
	Phone      *string   `db:"phone"`
	Username   *string   `db:"username"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
