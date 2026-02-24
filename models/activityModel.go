package models

import "time"

type Activity struct {
	ID          int       `db:"id" json:"id"`
	UserID      int       `db:"user_id" json:"user_id"`
	Action      string    `db:"activity_type" json:"action"`
	Description string    `db:"description" json:"description"`
	Page        string    `db:"page" json:"page"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
