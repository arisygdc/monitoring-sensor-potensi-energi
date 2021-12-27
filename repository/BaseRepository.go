package repository

import "monitoring-potensi-energi/database"

type Repository struct {
	Database database.DB
}

func New(db database.DB) (repo Repository) {
	repo = Repository{
		Database: db,
	}
	return
}
