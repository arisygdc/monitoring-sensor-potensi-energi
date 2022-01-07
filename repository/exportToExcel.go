package repository

import (
	"context"
	"monitoring-potensi-energi/database/postgres"
)

func (r Repository) GetAllValueSensor(ctx context.Context, idSensor int32) ([]postgres.GetAllValueSensorRow, error) {
	data, err := r.Database.Queries.GetAllValueSensor(ctx, idSensor)
	return data, err
}
