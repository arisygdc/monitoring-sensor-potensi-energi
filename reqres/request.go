package reqres

type SetupRequest struct {
	Sensor   Sensor   `json:"sensor" binding:"required"`
	Location Location `json:"lokasi" binding:"required"`
}

type InputValue struct {
	IDSensor int32   `json:"id_sensor" binding:"required"`
	Data     float64 `json:"data" binding:"required"`
}

type Sensor struct {
	TipeSensor string `json:"tipe_sensor" binding:"required"`
	Identity   string `json:"identity" binding:"required"`
}

type Location struct {
	Nama      string `json:"nama" binding:"required"`
	Provinsi  string `json:"provinsi" binding:"required"`
	Kecamatan string `json:"kecamatan" binding:"required"`
	Desa      string `json:"desa" binding:"required"`
}
