package gateway

import (
	"database/sql"
	"share-report/user/entity"

	"github.com/jmoiron/sqlx"
)

type organizationHandler struct {
	db *sqlx.DB
}

func NewOrganizationHandler(db *sqlx.DB) OrganizationHandler {
	return &organizationHandler{db}
}

func (oh *organizationHandler) GetOrganizationByID(id int64) (*entity.Organization, error) {
	var organization entity.Organization
	err := oh.db.Get(&organization, `
		SELECT * FROM organization WHERE id = ? AND valid = 1
	`, id)

	if err != nil && err == sql.ErrNoRows {
		return nil, entity.ErrOrganizationNotExist
	}
	
	if err != nil {
		return nil, err
	}
	return &organization, nil
}

func (oh *organizationHandler) CreateOrganization(o *entity.Organization) error {
	tx := oh.db.MustBegin()
	result := tx.MustExec(`
		INSERT INTO organization (user_id, name)
		VALUES (?, ?)
	`, o.UserID, o.Name)
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

func (oh *organizationHandler) UpdateOrganization(o *entity.Organization) error {
	tx := oh.db.MustBegin()
	tx.MustExec(`
		UPDATE organization SET user_id = ?, name = ?
		WHERE id = ?
	`, o.UserID, o.Name, o.ID)
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}