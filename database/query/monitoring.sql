-- name: AddTipeSensor :exec
INSERT INTO tipe_sensor (tipe, satuan) VALUES ($1, $2);

-- name: AddMonLocation :exec
INSERT INTO monitoring_location (provinsi, kecamatan, desa) VALUES ($1, $2, $3);

-- name: AddSensor :one
INSERT INTO sensors (tipe_sensor_id, mon_loc_id, status, ditempatkan_pada) VALUES ($1, $2, $3, $4) RETURNING id;

-- name: InputValueSensor :exec
INSERT INTO value_sensor (sensor_id, data, dibuat_pada) VALUES ($1, $2, $3);

-- name: GetTipeSensor :one
SELECT id FROM tipe_sensor WHERE tipe = $1;

-- name: GetMonitoringLocation :one
SELECT * FROM monitoring_location WHERE provinsi = $1 AND kecamatan = $2 AND desa = $3;

-- name: GetAllSensorByLocationID :many
SELECT * FROM sensors WHERE mon_loc_id = $1;

-- name: GetAllInSensorBetweenDate :many
SELECT * FROM value_sensor WHERE dibuat_pada BETWEEN $1 AND $2;

-- name: GetAllSensorOnStatus :many
SELECT s.id, s.ditempatkan_pada, MAX(vs.dibuat_pada) as last_update FROM sensors s
RIGHT JOIN value_sensor vs ON vs.sensor_id = s.id
WHERE status = $1 GROUP BY vs.dibuat_pada; 

-- name: UpdateStatusSensor :exec
UPDATE sensors SET status = $1 WHERE id = $2;
