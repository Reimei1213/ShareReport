package usecase

import (
	"share-report/user/adapter/gateway"
	"share-report/user/adapter/controller"
	"share-report/user/adapter/presenter"

	"github.com/jmoiron/sqlx"
)

type UserService struct {
	dh gateway.DatabaseHandler
	iph controller.InputPortHandler
	oph presenter.OutputPortHandler
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{
		dh: gateway.NewDatabaseHandler(db),
		iph: controller.NewInputPortHandler(),
		oph: presenter.NewOutputPortHandler(),
	}
}
