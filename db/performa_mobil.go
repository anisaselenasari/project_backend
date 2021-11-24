package db

type Performa struct {
	ID           int    `json:"id" form:"id"`
	Tenaga_Mobil string `json:"tenaga_mobil" form:"tenaga_mobil"`
	Torsi_Mobil  string `json:"torsi_mobil" form:"jml_mobil"`
}
