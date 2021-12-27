package database

import (
	"database/sql"
	"monitoring-potensi-energi/config"
	"monitoring-potensi-energi/database/postgres"

	_ "github.com/lib/pq"
)

type DB struct {
	SQLConn *sql.DB
	Queries *postgres.Queries
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
