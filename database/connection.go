package database

import (
	"database/sql"
	"monitoring-potensi-energi/config"

	_ "github.com/lib/pq"
)

type DB struct {
	SQLConn *sql.DB
}

func NewPostgres(env config.Environment) (postgreDB DB, err error) {
	SQLConn, err := sql.Open(env.DBDriver, env.DBSource)
	if err != nil {
		return
	}

	postgreDB = DB{
		SQLConn: SQLConn,
	}
	return
}
