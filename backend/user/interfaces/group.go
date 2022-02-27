package interfaces

import "github.com/jmoiron/sqlx"

type gropuHandler struct {
	db *sqlx.DB
}

func NewGroupHandler(db *sqlx.DB) GroupHandler {
	return gropuHandler{db}
}