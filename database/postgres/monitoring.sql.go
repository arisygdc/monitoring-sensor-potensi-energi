// Code generated by sqlc. DO NOT EDIT.
// source: monitoring.sql

package postgres

import (
	"context"
	"time"
)

const addInformasiSensor = `-- name: AddInformasiSensor :exec
INSERT INTO informasi_sensor (status, identity) VALUES ($1, $2)
`

type AddInformasiSensorParams struct {
	Status   bool   `json:"status"`
	Identity string `json:"identity"`
}

func (q *Queries) AddInformasiSensor(ctx context.Context, arg AddInformasiSensorParams) error {
	_, err := q.db.ExecContext(ctx, addInformasiSensor, arg.Status, arg.Identity)
	return err
}

const addMonLocation = `-- name: AddMonLocation :exec
INSERT INTO monitoring_location (nama, provinsi, kecamatan, desa) VALUES ($1, $2, $3, $4)
`

type AddMonLocationParams struct {
	Nama      string `json:"nama"`
	Provinsi  string `json:"provinsi"`
	Kecamatan string `json:"kecamatan"`
	Desa      string `json:"desa"`
}

func (q *Queries) AddMonLocation(ctx context.Context, arg AddMonLocationParams) error {
	_, err := q.db.ExecContext(ctx, addMonLocation,
		arg.Nama,
		arg.Provinsi,
		arg.Kecamatan,
		arg.Desa,
	)
	return err
}

const addSensor = `-- name: AddSensor :exec
INSERT INTO sensors (tipe_sensor_id, inf_sensor_id, mon_loc_id, ditempatkan_pada) VALUES ($1, $2, $3, $4)
`

type AddSensorParams struct {
	TipeSensorID    int32     `json:"tipe_sensor_id"`
	InfSensorID     int32     `json:"inf_sensor_id"`
	MonLocID        int32     `json:"mon_loc_id"`
	DitempatkanPada time.Time `json:"ditempatkan_pada"`
}

func (q *Queries) AddSensor(ctx context.Context, arg AddSensorParams) error {
	_, err := q.db.ExecContext(ctx, addSensor,
		arg.TipeSensorID,
		arg.InfSensorID,
		arg.MonLocID,
		arg.DitempatkanPada,
	)
	return err
}

const addTipeSensor = `-- name: AddTipeSensor :exec
INSERT INTO tipe_sensor (tipe) VALUES ($1)
`

func (q *Queries) AddTipeSensor(ctx context.Context, tipe string) error {
	_, err := q.db.ExecContext(ctx, addTipeSensor, tipe)
	return err
}

const getAllInSensorBetweenDate = `-- name: GetAllInSensorBetweenDate :many
SELECT id, sensor_id, data, dibuat_pada FROM value_sensor WHERE dibuat_pada BETWEEN $1 AND $2
`

type GetAllInSensorBetweenDateParams struct {
	DibuatPada   time.Time `json:"dibuat_pada"`
	DibuatPada_2 time.Time `json:"dibuat_pada_2"`
}

func (q *Queries) GetAllInSensorBetweenDate(ctx context.Context, arg GetAllInSensorBetweenDateParams) ([]ValueSensor, error) {
	rows, err := q.db.QueryContext(ctx, getAllInSensorBetweenDate, arg.DibuatPada, arg.DibuatPada_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ValueSensor
	for rows.Next() {
		var i ValueSensor
		if err := rows.Scan(
			&i.ID,
			&i.SensorID,
			&i.Data,
			&i.DibuatPada,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllSensorByIdentity = `-- name: GetAllSensorByIdentity :many
SELECT s.id, inf_sensor_id, tipe_sensor_id, mon_loc_id, ditempatkan_pada, si.id, status, identity FROM sensors s INNER JOIN informasi_sensor si ON si.id = s.inf_sensor_id WHERE si.identity
`

type GetAllSensorByIdentityRow struct {
	ID              int64     `json:"id"`
	InfSensorID     int32     `json:"inf_sensor_id"`
	TipeSensorID    int32     `json:"tipe_sensor_id"`
	MonLocID        int32     `json:"mon_loc_id"`
	DitempatkanPada time.Time `json:"ditempatkan_pada"`
	ID_2            int32     `json:"id_2"`
	Status          bool      `json:"status"`
	Identity        string    `json:"identity"`
}

func (q *Queries) GetAllSensorByIdentity(ctx context.Context) ([]GetAllSensorByIdentityRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllSensorByIdentity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllSensorByIdentityRow
	for rows.Next() {
		var i GetAllSensorByIdentityRow
		if err := rows.Scan(
			&i.ID,
			&i.InfSensorID,
			&i.TipeSensorID,
			&i.MonLocID,
			&i.DitempatkanPada,
			&i.ID_2,
			&i.Status,
			&i.Identity,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllSensorByLocationID = `-- name: GetAllSensorByLocationID :many
SELECT id, inf_sensor_id, tipe_sensor_id, mon_loc_id, ditempatkan_pada FROM sensors WHERE mon_loc_id = $1
`

func (q *Queries) GetAllSensorByLocationID(ctx context.Context, monLocID int32) ([]Sensor, error) {
	rows, err := q.db.QueryContext(ctx, getAllSensorByLocationID, monLocID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Sensor
	for rows.Next() {
		var i Sensor
		if err := rows.Scan(
			&i.ID,
			&i.InfSensorID,
			&i.TipeSensorID,
			&i.MonLocID,
			&i.DitempatkanPada,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInfSensor = `-- name: GetInfSensor :one
SELECT id, status, identity FROM informasi_sensor WHERE identity = $1
`

func (q *Queries) GetInfSensor(ctx context.Context, identity string) (InformasiSensor, error) {
	row := q.db.QueryRowContext(ctx, getInfSensor, identity)
	var i InformasiSensor
	err := row.Scan(&i.ID, &i.Status, &i.Identity)
	return i, err
}

const getMonitoringLocation = `-- name: GetMonitoringLocation :one
SELECT id, nama, provinsi, kecamatan, desa FROM monitoring_location WHERE  nama = $1 AND provinsi = $2 AND kecamatan = $3 AND desa = $4
`

type GetMonitoringLocationParams struct {
	Nama      string `json:"nama"`
	Provinsi  string `json:"provinsi"`
	Kecamatan string `json:"kecamatan"`
	Desa      string `json:"desa"`
}

func (q *Queries) GetMonitoringLocation(ctx context.Context, arg GetMonitoringLocationParams) (MonitoringLocation, error) {
	row := q.db.QueryRowContext(ctx, getMonitoringLocation,
		arg.Nama,
		arg.Provinsi,
		arg.Kecamatan,
		arg.Desa,
	)
	var i MonitoringLocation
	err := row.Scan(
		&i.ID,
		&i.Nama,
		&i.Provinsi,
		&i.Kecamatan,
		&i.Desa,
	)
	return i, err
}

const getTipeSensor = `-- name: GetTipeSensor :one
SELECT id, tipe FROM tipe_sensor WHERE tipe = $1
`

func (q *Queries) GetTipeSensor(ctx context.Context, tipe string) (TipeSensor, error) {
	row := q.db.QueryRowContext(ctx, getTipeSensor, tipe)
	var i TipeSensor
	err := row.Scan(&i.ID, &i.Tipe)
	return i, err
}

const inputValueSensor = `-- name: InputValueSensor :exec
INSERT INTO value_sensor (sensor_id, data) VALUES ($1, $2)
`

type InputValueSensorParams struct {
	SensorID int32   `json:"sensor_id"`
	Data     float64 `json:"data"`
}

func (q *Queries) InputValueSensor(ctx context.Context, arg InputValueSensorParams) error {
	_, err := q.db.ExecContext(ctx, inputValueSensor, arg.SensorID, arg.Data)
	return err
}
