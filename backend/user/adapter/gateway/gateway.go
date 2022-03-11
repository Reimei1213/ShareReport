package gateway

import (
	"share-report/user/entity"

	"github.com/jmoiron/sqlx"
)

type DatabaseHandler interface {
	UserHandler
	OrganizationUserHandler
	OrganizationHandler
}

type UserHandler interface {
	GetUserByID(id string) (*entity.User, error)
	CreateUser(u *entity.User) error
	UpdateUser(u *entity.User) error
	DeleteUserById(id string) error
}

type OrganizationUserHandler interface {
}

type OrganizationHandler interface {
	GetOrganizationByID(id int64) (*entity.Organization, error)
	CreateOrganization(o *entity.Organization) error
	UpdateOrganization(o *entity.Organization) error
}

type databaseHandler struct {
	UserHandler
	OrganizationUserHandler
	OrganizationHandler
}

func NewDatabaseHandler(db *sqlx.DB) DatabaseHandler {
	uh := NewUserHandler(db)
	ouh := NewOrganizationUserHandler(db)
	oh := NewOrganizationHandler(db)
	return &databaseHandler{uh, ouh, oh}
}
