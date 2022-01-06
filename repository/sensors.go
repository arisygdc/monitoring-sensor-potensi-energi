package repository

import (
	"context"
	"monitoring-potensi-energi/database/postgres"
)

func (r Repository) GetAllSensor(ctx context.Context) ([]postgres.GetSensorsRow, error) {
	return r.Database.Queries.GetSensors(ctx)
}

func (r Repository) GetMetrics(ctx context.Context, id int32) ([]postgres.GetValueSensorRow, error) {
	return r.Database.Queries.GetValueSensor(ctx, id)
}
