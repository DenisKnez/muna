package util

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	//database driver
	_ "github.com/jackc/pgx/v4/stdlib"
	ini "gopkg.in/ini.v1"
)

//GetDatabaseConnection get the database connection with the provided config
func GetDatabaseConnection(config *ini.File) *sql.DB {

	connString := config.Section("database").Key("connection").String()

	var err error
	Db, err := sql.Open("pgx", connString)

	if err != nil {
		fmt.Println(err)
		panic("Cannot connect to database")
	}

	return Db
}

//NewUUID generates a new uuid
func NewUUID() uuid.UUID {
	return uuid.New()
}
