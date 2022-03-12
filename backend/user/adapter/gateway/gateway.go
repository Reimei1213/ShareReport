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
	GetUserListByOrganizationId(organization_id int64) ([]*entity.User, error)
	CreateUser(u *entity.User) error
	UpdateUser(u *entity.User) error
	DeleteUserByID(id string) error
}

type OrganizationUserHandler interface {
	GetOrganizationUserByUserIDAndOrganizationID(user_id string, organization_id int64) (*entity.OrganizationUser, error)
	CreateOrganizationUser(ou *entity.OrganizationUser) error
	DeleteOrganizationUserByUserID(user_id string) error
	DeleteOrganizationUserByOrganizationID(organization_id int64) error
}

type OrganizationHandler interface {
	GetOrganizationByID(id int64) (*entity.Organization, error)
	GetOrganizationListByUserId(user_id string) ([]*entity.Organization, error)
	CreateOrganization(o *entity.Organization) error
	UpdateOrganization(o *entity.Organization) error
	DeleteOrganizationByID(id int64) error
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
