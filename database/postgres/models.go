// Code generated by sqlc. DO NOT EDIT.

package postgres

import (
	"time"
)

type InformasiSensor struct {
	ID       int32  `json:"id"`
	Status   bool   `json:"status"`
	Identity string `json:"identity"`
}

type MonitoringLocation struct {
	ID        int32  `json:"id"`
	Nama      string `json:"nama"`
	Provinsi  string `json:"provinsi"`
	Kecamatan string `json:"kecamatan"`
	Desa      string `json:"desa"`
}

type Sensor struct {
	ID              int64     `json:"id"`
	InfSensorID     int32     `json:"inf_sensor_id"`
	TipeSensorID    int32     `json:"tipe_sensor_id"`
	MonLocID        int32     `json:"mon_loc_id"`
	DitempatkanPada time.Time `json:"ditempatkan_pada"`
}

type TipeSensor struct {
	ID   int32  `json:"id"`
	Tipe string `json:"tipe"`
}

type ValueSensor struct {
	ID         int64     `json:"id"`
	SensorID   int32     `json:"sensor_id"`
	Data       float64   `json:"data"`
	DibuatPada time.Time `json:"dibuat_pada"`
}