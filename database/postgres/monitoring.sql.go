// Code generated by sqlc. DO NOT EDIT.
// source: monitoring.sql

package postgres

import (
	"context"
	"database/sql"
	"time"
)

const addMonLocation = `-- name: AddMonLocation :exec
INSERT INTO monitoring_location (provinsi, kecamatan, desa) VALUES ($1, $2, $3)
`

type AddMonLocationParams struct {
	Provinsi  string `json:"provinsi"`
	Kecamatan string `json:"kecamatan"`
	Desa      string `json:"desa"`
}

func (q *Queries) AddMonLocation(ctx context.Context, arg AddMonLocationParams) error {
	_, err := q.db.ExecContext(ctx, addMonLocation, arg.Provinsi, arg.Kecamatan, arg.Desa)
	return err
}

const addSensor = `-- name: AddSensor :one
INSERT INTO sensors (tipe_sensor_id, mon_loc_id, status, ditempatkan_pada) VALUES ($1, $2, $3, $4) RETURNING id
`

type AddSensorParams struct {
	TipeSensorID    int32     `json:"tipe_sensor_id"`
	MonLocID        int32     `json:"mon_loc_id"`
	Status          bool      `json:"status"`
	DitempatkanPada time.Time `json:"ditempatkan_pada"`
}

func (q *Queries) AddSensor(ctx context.Context, arg AddSensorParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, addSensor,
		arg.TipeSensorID,
		arg.MonLocID,
		arg.Status,
		arg.DitempatkanPada,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const addTipeSensor = `-- name: AddTipeSensor :exec
INSERT INTO tipe_sensor (tipe, satuan) VALUES ($1, $2)
`

type AddTipeSensorParams struct {
	Tipe   string `json:"tipe"`
	Satuan string `json:"satuan"`
}

func (q *Queries) AddTipeSensor(ctx context.Context, arg AddTipeSensorParams) error {
	_, err := q.db.ExecContext(ctx, addTipeSensor, arg.Tipe, arg.Satuan)
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

const getAllSensorByLocationID = `-- name: GetAllSensorByLocationID :many
SELECT id, tipe_sensor_id, mon_loc_id, status, ditempatkan_pada FROM sensors WHERE mon_loc_id = $1
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
			&i.TipeSensorID,
			&i.MonLocID,
			&i.Status,
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

const getAllSensorOnStatus = `-- name: GetAllSensorOnStatus :many
SELECT s.id, s.ditempatkan_pada, MAX(vs.dibuat_pada) as terakhir_update FROM sensors s
LEFT JOIN value_sensor vs ON vs.sensor_id = s.id
WHERE s.status = $1 group by s.id order by s.id asc
`

type GetAllSensorOnStatusRow struct {
	ID              int64       `json:"id"`
	DitempatkanPada time.Time   `json:"ditempatkan_pada"`
	TerakhirUpdate  interface{} `json:"terakhir_update"`
}

func (q *Queries) GetAllSensorOnStatus(ctx context.Context, status bool) ([]GetAllSensorOnStatusRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllSensorOnStatus, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllSensorOnStatusRow
	for rows.Next() {
		var i GetAllSensorOnStatusRow
		if err := rows.Scan(&i.ID, &i.DitempatkanPada, &i.TerakhirUpdate); err != nil {
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

const getMonitoringLocation = `-- name: GetMonitoringLocation :one
SELECT id, provinsi, kecamatan, desa FROM monitoring_location WHERE provinsi = $1 AND kecamatan = $2 AND desa = $3
`

type GetMonitoringLocationParams struct {
	Provinsi  string `json:"provinsi"`
	Kecamatan string `json:"kecamatan"`
	Desa      string `json:"desa"`
}

func (q *Queries) GetMonitoringLocation(ctx context.Context, arg GetMonitoringLocationParams) (MonitoringLocation, error) {
	row := q.db.QueryRowContext(ctx, getMonitoringLocation, arg.Provinsi, arg.Kecamatan, arg.Desa)
	var i MonitoringLocation
	err := row.Scan(
		&i.ID,
		&i.Provinsi,
		&i.Kecamatan,
		&i.Desa,
	)
	return i, err
}

const getSensors = `-- name: GetSensors :many
SELECT s.id, ts.tipe, ml.provinsi, ml.kecamatan, ml.desa, s.ditempatkan_pada, s.status FROM sensors s 
RIGHT JOIN tipe_sensor ts ON s.tipe_sensor_id = ts.id
RIGHT JOIN monitoring_location ml ON s.mon_loc_id = ml.id
LIMIT 30
`

type GetSensorsRow struct {
	ID              sql.NullInt64 `json:"id"`
	Tipe            string        `json:"tipe"`
	Provinsi        string        `json:"provinsi"`
	Kecamatan       string        `json:"kecamatan"`
	Desa            string        `json:"desa"`
	DitempatkanPada sql.NullTime  `json:"ditempatkan_pada"`
	Status          sql.NullBool  `json:"status"`
}

func (q *Queries) GetSensors(ctx context.Context) ([]GetSensorsRow, error) {
	rows, err := q.db.QueryContext(ctx, getSensors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSensorsRow
	for rows.Next() {
		var i GetSensorsRow
		if err := rows.Scan(
			&i.ID,
			&i.Tipe,
			&i.Provinsi,
			&i.Kecamatan,
			&i.Desa,
			&i.DitempatkanPada,
			&i.Status,
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

const getSensorsOnStatus = `-- name: GetSensorsOnStatus :many
SELECT s.id, tipe_sensor_id, mon_loc_id, status, ditempatkan_pada, ts.id, tipe, satuan, ml.id, provinsi, kecamatan, desa FROM sensors s 
RIGHT JOIN tipe_sensor ts ON s.tipe_sensor_id = ts.id
RIGHT JOIN monitoring_location ml ON s.mon_loc_id = ml.id
WHERE s.status = $1 LIMIT 30
`

type GetSensorsOnStatusRow struct {
	ID              int64     `json:"id"`
	TipeSensorID    int32     `json:"tipe_sensor_id"`
	MonLocID        int32     `json:"mon_loc_id"`
	Status          bool      `json:"status"`
	DitempatkanPada time.Time `json:"ditempatkan_pada"`
	ID_2            int32     `json:"id_2"`
	Tipe            string    `json:"tipe"`
	Satuan          string    `json:"satuan"`
	ID_3            int32     `json:"id_3"`
	Provinsi        string    `json:"provinsi"`
	Kecamatan       string    `json:"kecamatan"`
	Desa            string    `json:"desa"`
}

func (q *Queries) GetSensorsOnStatus(ctx context.Context, status bool) ([]GetSensorsOnStatusRow, error) {
	rows, err := q.db.QueryContext(ctx, getSensorsOnStatus, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSensorsOnStatusRow
	for rows.Next() {
		var i GetSensorsOnStatusRow
		if err := rows.Scan(
			&i.ID,
			&i.TipeSensorID,
			&i.MonLocID,
			&i.Status,
			&i.DitempatkanPada,
			&i.ID_2,
			&i.Tipe,
			&i.Satuan,
			&i.ID_3,
			&i.Provinsi,
			&i.Kecamatan,
			&i.Desa,
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

const getTipeSensor = `-- name: GetTipeSensor :one
SELECT id FROM tipe_sensor WHERE tipe = $1
`

func (q *Queries) GetTipeSensor(ctx context.Context, tipe string) (int32, error) {
	row := q.db.QueryRowContext(ctx, getTipeSensor, tipe)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const inputValueSensor = `-- name: InputValueSensor :exec
INSERT INTO value_sensor (sensor_id, data, dibuat_pada) VALUES ($1, $2, $3)
`

type InputValueSensorParams struct {
	SensorID   int32     `json:"sensor_id"`
	Data       float64   `json:"data"`
	DibuatPada time.Time `json:"dibuat_pada"`
}

func (q *Queries) InputValueSensor(ctx context.Context, arg InputValueSensorParams) error {
	_, err := q.db.ExecContext(ctx, inputValueSensor, arg.SensorID, arg.Data, arg.DibuatPada)
	return err
}

const updateStatusSensor = `-- name: UpdateStatusSensor :exec
UPDATE sensors SET status = $1 WHERE id = $2
`

type UpdateStatusSensorParams struct {
	Status bool  `json:"status"`
	ID     int64 `json:"id"`
}

func (q *Queries) UpdateStatusSensor(ctx context.Context, arg UpdateStatusSensorParams) error {
	_, err := q.db.ExecContext(ctx, updateStatusSensor, arg.Status, arg.ID)
	return err
}
