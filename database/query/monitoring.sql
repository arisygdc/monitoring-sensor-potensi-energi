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

-- name: GetAllValueSensor :many
SELECT data, dibuat_pada FROM value_sensor WHERE sensor_id = $1;

-- name: GetAllSensorOnStatus :many
SELECT s.id, s.ditempatkan_pada, MAX(vs.dibuat_pada) as terakhir_update FROM sensors s
LEFT JOIN value_sensor vs ON vs.sensor_id = s.id
WHERE s.status = $1 group by s.id order by s.id asc;

-- name: UpdateStatusSensor :exec
UPDATE sensors SET status = $1 WHERE id = $2;

-- name: GetValueSensor :many 
SELECT data, dibuat_pada FROM value_sensor WHERE sensor_id = $1 ORDER BY id DESC LIMIT 30;

-- name: GetSensors :many
SELECT s.id, ts.tipe, ml.provinsi, ml.kecamatan, ml.desa, s.ditempatkan_pada, s.status FROM sensors s 
RIGHT JOIN tipe_sensor ts ON s.tipe_sensor_id = ts.id
RIGHT JOIN monitoring_location ml ON s.mon_loc_id = ml.id
LIMIT 30;