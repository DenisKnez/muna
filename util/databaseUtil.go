package util

import (
	"database/sql"

	"github.com/google/uuid"
)

//Db Database connection
var Db *sql.DB

func init() {

	var err error
	Db, err = sql.Open("pqx", "postgres://postgres:rootPassword@localhost:5432/muna?sslmode=disable")

	if err != nil {
		panic("Cannot connect to database")
	}

}

//NewUUID generates a new uuid
func NewUUID() uuid.UUID {
	return uuid.New()
}
