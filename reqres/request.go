package reqres

type SetupRequest struct {
	TipeSensor string `json:"tipe_sensor"`
	Identity   string `json:"identity"`
	NamaLokasi string `json:"nama_lokasi"`
	Provinsi   string `json:"provinsi"`
	Kecamatan  string `json:"kecamatan"`
	Desa       string `json:"desa"`
}