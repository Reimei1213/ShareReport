package gateway

import "github.com/jmoiron/sqlx"

type organizationHandler struct {
	db *sqlx.DB
}

func NewOrganizationHandler(db *sqlx.DB) OrganizationHandler {
	return &organizationHandler{db}
}
