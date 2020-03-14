package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDb() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "rlv"
	dbPass := "hola"
	dbName := "weero"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
