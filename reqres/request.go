package reqres

type SetupRequest struct {
	TipeSensor string `JSON:"tipe_sensor"`
	Identity   string `JSON:"identity"`
	NamaLokasi string `JSON:"nama_lokasi"`
	Provinsi   string `JSON:"provinsi"`
	Kecamatan  string `JSON:"kecamatan"`
	Desa       string `JSON:"desa"`
}
