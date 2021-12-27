package repository

import (
	"context"
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

func (r Repository) transaction(ctx context.Context, query func(q *postgres.Queries) error) error {
	tx, err := r.Database.SQLConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	txQueries := r.Database.Queries.WithTx(tx)
	err = query(txQueries)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit()
}
