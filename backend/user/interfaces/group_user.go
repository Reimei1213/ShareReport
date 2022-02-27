package interfaces

import "github.com/jmoiron/sqlx"

type groupUserHandler struct {
	db *sqlx.DB
}

func NewGroupUserHandler(db *sqlx.DB) GroupUserHandler {
	return groupUserHandler{db}
}