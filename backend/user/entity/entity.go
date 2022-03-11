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

type OrganizationUser struct {
	ID             int64     `db:"id"`
	UserID         string    `db:"user_id"`
	OrganizationID int64     `db:"organization_id"`
	Valid          bool      `db:"valid"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type Organization struct {
	ID        int64     `db:"id"`
	UserID    string    `db:"user_id"`
	Name      string    `db:"name"`
	Valid     bool      `db:"valid"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
