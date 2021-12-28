package repository

import (
	"context"
	"database/sql"
	"monitoring-potensi-energi/database"
	"monitoring-potensi-energi/database/postgres"
)

type Repository struct {
	Database database.DB
}

func New(db database.DB) (repo Repository) {
	repo = Repository{
		Database: db,
	}
	return
}

func (r Repository) CreateTX(ctx context.Context) (tx *sql.Tx, txQueries *postgres.Queries, err error) {
	tx, err = r.Database.SQLConn.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	txQueries = r.Database.Queries.WithTx(tx)
	return
}

func (r Repository) execTX(ctx context.Context, query func(q *postgres.Queries) error) error {
	tx, txQueries, err := r.CreateTX(ctx)
	if err != nil {
		return err
	}

	err = query(txQueries)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit()
}
