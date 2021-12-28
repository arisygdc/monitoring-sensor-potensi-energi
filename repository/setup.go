package repository

import (
	"context"
	"errors"
	"monitoring-potensi-energi/database/postgres"
	"monitoring-potensi-energi/reqres"
	"time"
)

func (r Repository) PlaceSensor(ctx context.Context, req reqres.SetupRequest) (int64, error) {
	var idSensor int64
	tipe, err := r.Database.Queries.GetTipeSensor(ctx, req.Sensor.TipeSensor)
	if err != nil {
		return idSensor, errors.New("tipe sensor tidak ditemukan")
	}

	err = r.execTX(ctx, func(q *postgres.Queries) error {
		infSensor, err := r.setupInfSensor(ctx, q, req.Sensor.Identity)
		if err != nil {
			return err
		}

		monLocArg := postgres.GetMonitoringLocationParams{
			Nama:      req.Location.Nama,
			Provinsi:  req.Location.Provinsi,
			Kecamatan: req.Location.Kecamatan,
			Desa:      req.Location.Desa,
		}

		monLoc, err := r.setupLocation(ctx, q, monLocArg)
		if err != nil {
			return err
		}

		addSensor := postgres.AddSensorParams{
			TipeSensorID:    tipe.ID,
			InfSensorID:     infSensor.ID,
			MonLocID:        monLoc.ID,
			DitempatkanPada: time.Now().UTC(),
		}

		idSensor, err = q.AddSensor(ctx, addSensor)
		if err != nil {
			return err
		}

		return nil
	})

	return idSensor, err
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

func (r Repository) setupInfSensor(ctx context.Context, q *postgres.Queries, identity string) (inf postgres.InformasiSensor, err error) {
	inf, err = q.GetInfSensor(ctx, identity)
	if err != nil {
		if err = q.AddInformasiSensor(ctx, postgres.AddInformasiSensorParams{
			Status:   true,
			Identity: identity,
		}); err != nil {
			return
		}

		inf, err = q.GetInfSensor(ctx, identity)
	}
	return
}
