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
	tipe, err := r.Database.Queries.GetTipeSensor(ctx, req.TipeSensor)
	if err != nil {
		return idSensor, errors.New("tipe sensor tidak ditemukan")
	}

	err = r.transaction(ctx, func(q *postgres.Queries) error {
		infSensor, err := q.GetInfSensor(ctx, req.Identity)
		if err != nil {
			if err := q.AddInformasiSensor(ctx, postgres.AddInformasiSensorParams{
				Status:   true,
				Identity: req.Identity,
			}); err != nil {
				return err
			}

			infSensor, err = q.GetInfSensor(ctx, req.Identity)
			if err != nil {
				return err
			}
		}

		monLocArg := postgres.GetMonitoringLocationParams{
			Nama:      req.NamaLokasi,
			Provinsi:  req.Provinsi,
			Kecamatan: req.Kecamatan,
			Desa:      req.Desa,
		}

		monLoc, err := q.GetMonitoringLocation(ctx, monLocArg)
		if err != nil {
			if err := q.AddMonLocation(ctx, postgres.AddMonLocationParams{
				Nama:      req.NamaLokasi,
				Provinsi:  req.Provinsi,
				Kecamatan: req.Kecamatan,
				Desa:      req.Desa,
			}); err != nil {
				return err
			}

			monLoc, err = q.GetMonitoringLocation(ctx, monLocArg)
			if err != nil {
				return err
			}
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
