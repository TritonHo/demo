package handler

import (
	"database/sql"
)

var (
	db *sql.DB
)

func Init(database *sql.DB) {
	db = database
}
