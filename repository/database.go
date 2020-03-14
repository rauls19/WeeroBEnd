package repository

import (
	"database/sql"
)

//Connection is a object to save the sql.DB reference
type Connection struct {
	db *sql.DB
}

//InstanceDB is the function which allows save the instance in an object and return it
func InstanceDB(db *sql.DB) *Connection {
	return &Connection{db}
}
