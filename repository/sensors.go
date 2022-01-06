package repository

import (
	"context"
	"monitoring-potensi-energi/database/postgres"
)

func (r Repository) GetAllActiveSensor(ctx context.Context) ([]postgres.GetSensorsOnStatusRow, error) {
	return r.Database.Queries.GetSensorsOnStatus(ctx, true)
}

func (r Repository) GetAllSensor(ctx context.Context) ([]postgres.GetSensorsRow, error) {
	return r.Database.Queries.GetSensors(ctx)
}
