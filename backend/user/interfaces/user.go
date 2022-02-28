package interfaces

import (
	"share-report/user/entity"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	db *sqlx.DB
}

func NewUserHandler(db *sqlx.DB) UserHandler {
	return &userHandler{db}
}

func (uh *userHandler) CreateUser(u *entity.User) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}

	tx := uh.db.MustBegin()
	result := tx.MustExec(`
		INSERT INTO user (id, name, email, password)
		VALUES (?, ?, ?, ?)
	`, uuid.String(), u.Name, u.Email, string(hashedPassword))
	_, err = result.LastInsertId()
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