package repository

import (
	"context"
	"monitoring-potensi-energi/database/postgres"
	"monitoring-potensi-energi/reqres"
	"time"
)

func (r Repository) InputValue(ctx context.Context, req reqres.InputValue) error {
	err := r.Database.Queries.InputValueSensor(ctx, postgres.InputValueSensorParams{
		SensorID:   req.IDSensor,
		Data:       req.Data,
		DibuatPada: time.Now().UTC(),
	})
	return err
}
