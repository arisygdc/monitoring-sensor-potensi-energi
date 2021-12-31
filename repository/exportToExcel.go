package repository

import (
	"context"
	"monitoring-potensi-energi/database/postgres"
	"time"
)

func (r Repository) GetMonitoringDataBetween(ctx context.Context, startTime time.Time, until time.Time) ([]postgres.ValueSensor, error) {
	data, err := r.Database.Queries.GetAllInSensorBetweenDate(ctx, postgres.GetAllInSensorBetweenDateParams{
		DibuatPada:   startTime,
		DibuatPada_2: until,
	})
	return []postgres.ValueSensor(data), err
}
