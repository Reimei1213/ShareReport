package gateway

import "github.com/jmoiron/sqlx"

type organizationUserHandler struct {
	db *sqlx.DB
}

func NewOrganizationUserHandler(db *sqlx.DB) OrganizationUserHandler {
	return &organizationUserHandler{db}
}
