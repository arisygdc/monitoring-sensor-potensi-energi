package reqres

type SetupRequest struct {
	Sensors  []string `json:"sensors" binding:"required"`
	Location Location `json:"lokasi" binding:"required"`
}

type InputValue struct {
	IDSensor int32   `json:"id_sensor" binding:"required"`
	Data     float64 `json:"data" binding:"required"`
}

type Location struct {
	Provinsi  string `json:"provinsi" binding:"required"`
	Kecamatan string `json:"kecamatan" binding:"required"`
	Desa      string `json:"desa" binding:"required"`
}
