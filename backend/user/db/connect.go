package db

import (
	"os"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Connect() (*sqlx.DB, error) {
	u := os.Getenv("USER_DB_USER")
	pw := os.Getenv("USER_DB_PASSWORD")
	h := os.Getenv("USER_DB_HOST")
	p := os.Getenv("USER_DB_PORT")
	dn := os.Getenv("DB_NAME")
	dbms := os.Getenv("USER_DB_DBMS")

	dbconf := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", u, pw, h, p, dn)
	db, err := sqlx.Connect(dbms, dbconf)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
