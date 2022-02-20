package usecase

import (
	"share-report/user/interfaces"

	"github.com/jmoiron/sqlx"
)

type UserService struct {
	dh interfaces.DatabaseHandler
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{
		dh: interfaces.NewDatabaseHandler(db),
	}
}
