package gateway

import (
	"database/sql"
	"share-report/user/entity"

	"github.com/jmoiron/sqlx"
)

type organizationUserHandler struct {
	db *sqlx.DB
}

func NewOrganizationUserHandler(db *sqlx.DB) OrganizationUserHandler {
	return &organizationUserHandler{db}
}

func (ouh *organizationUserHandler) GetOrganizationByUserIDAndOrganizationID(user_id string, organization_id int64) (*entity.OrganizationUser, error) {
	var organizationUser entity.OrganizationUser
	err := ouh.db.Get(&organizationUser, `
		SELECT * FROM organization_user WHERE user_id = ? AND organization_id = ? AND valid = 1
	`, user_id, organization_id)

	if err != nil && err == sql.ErrNoRows {
		return nil, entity.ErrOrganizationUserNotExist
	}

	if err != nil {
		return nil, err
	}
	return &organizationUser, nil
}

func (ouh *organizationUserHandler) CreateOrganizationUser(ou *entity.OrganizationUser) error {
	tx := ouh.db.MustBegin()
	result := tx.MustExec(`
		INSERT INTO organization_user (user_id, organization_id)
		VALUES (?, ?)
	`, ou.UserID, ou.OrganizationID)
	_, err := result.LastInsertId()
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err2
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}