package repository

import (
	"context"
	"errors"
	"monitoring-potensi-energi/database/postgres"
	"monitoring-potensi-energi/reqres"
	"time"
)

func (r Repository) PlaceSensor(ctx context.Context, req reqres.SetupRequest) (sensor map[string]int32, err error) {
	sensor = make(map[string]int32)
	err = r.execTX(ctx, func(q *postgres.Queries) error {
		monLocArg := postgres.GetMonitoringLocationParams{
			Provinsi:  req.Location.Provinsi,
			Kecamatan: req.Location.Kecamatan,
			Desa:      req.Location.Desa,
		}

		monLoc, err := r.setupLocation(ctx, q, monLocArg)
		if err != nil {
			return err
		}

		for _, v := range req.Sensors {
			tipe, err := r.Database.Queries.GetTipeSensor(ctx, v)
			if err != nil {
				err = errors.New("tipe sensor tidak ditemukan")
				return err
			}

			addSensor := postgres.AddSensorParams{
				TipeSensorID:    tipe,
				MonLocID:        monLoc.ID,
				DitempatkanPada: time.Now().UTC(),
			}

			id, err := q.AddSensor(ctx, addSensor)
			if err != nil {
				return err
			}
			sensor[v] = int32(id)
		}

		return nil
	})

	return sensor, err
}

func (r Repository) setupLocation(ctx context.Context, q *postgres.Queries, param postgres.GetMonitoringLocationParams) (monLoc postgres.MonitoringLocation, err error) {
	monLoc, err = q.GetMonitoringLocation(ctx, param)
	if err != nil {
		if err = q.AddMonLocation(ctx, postgres.AddMonLocationParams(param)); err != nil {
			return
		}

		monLoc, err = q.GetMonitoringLocation(ctx, param)
	}
	return
}
