package entity

import "time"

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Valid     bool      `db:"valid"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type GroupUser struct {
	ID        int       `db:"id"`
	UserID    string    `db:"user_id"`
	GroupID   int       `db:"group_id"`
	Valid     bool      `db:"valid"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Group struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Valid     bool      `db:"valid"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
