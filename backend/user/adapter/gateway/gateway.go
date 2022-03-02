package gateway

import (
	"share-report/user/entity"

	"github.com/jmoiron/sqlx"
)

type DatabaseHandler interface {
	UserHandler
	GroupUserHandler
	GroupHandler
}

type UserHandler interface {
	GetUserByID(id string) (*entity.User, error)
	CreateUser(u *entity.User) error
	UpdateUser(u *entity.User) error
	DeleteUserById(id string) error
}

type GroupUserHandler interface {
}

type GroupHandler interface {
}

type databaseHandler struct {
	UserHandler
	GroupUserHandler
	GroupHandler
}

func NewDatabaseHandler(db *sqlx.DB) DatabaseHandler {
	uh := NewUserHandler(db)
	guh := NewGroupUserHandler(db)
	gh := NewGroupHandler(db)
	return &databaseHandler{uh, guh, gh}
}
