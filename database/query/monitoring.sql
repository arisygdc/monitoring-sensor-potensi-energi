-- name: AddTipeSensor :exec
INSERT INTO tipe_sensor (tipe) VALUES ($1);

-- name: AddInformasiSensor :exec
INSERT INTO informasi_sensor (status, identity) VALUES ($1, $2);

-- name: AddMonLocation :exec
INSERT INTO monitoring_location (nama, provinsi, kecamatan, desa) VALUES ($1, $2, $3, $4);

-- name: AddSensor :exec
INSERT INTO sensors (tipe_sensor_id, inf_sensor_id, mon_loc_id, ditempatkan_pada) VALUES ($1, $2, $3, $4);

-- name: InputValueSensor :exec
INSERT INTO value_sensor (sensor_id, data) VALUES ($1, $2);

-- name: GetTipeSensor :one
SELECT * FROM tipe_sensor WHERE tipe = $1;

-- name: GetInfSensor :one
SELECT * FROM informasi_sensor WHERE identity = $1;

-- name: GetMonitoringLocation :one
SELECT * FROM monitoring_location WHERE  nama = $1 AND provinsi = $2 AND kecamatan = $3 AND desa = $4;

-- name: GetAllSensorByLocationID :many
SELECT * FROM sensors WHERE mon_loc_id = $1;

-- name: GetAllSensorByIdentity :many
SELECT * FROM sensors s INNER JOIN informasi_sensor si ON si.id = s.inf_sensor_id WHERE si.identity;

-- name: GetAllInSensorBetweenDate :many
SELECT * FROM value_sensor WHERE dibuat_pada BETWEEN $1 AND $2;
