package interfaces

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
	CreateUser(u entity.User) error
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