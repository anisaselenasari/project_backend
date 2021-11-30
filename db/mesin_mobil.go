package db

type Mesin struct {
	ID                    int    `json:"id" form:"id"`
	TipeDrive             string `json:"tipedrive" form:"tipedrive"`
	Kapasitas_Tangki      string `json:"kapasitas_tangki" form:"kapasitas_tangki"`
	Jenis_BahanBakar      string `json:"jenis_bahanbakar" form:"jenis_bahanbakar"`
	Konfigurasi_Transmisi string `json:"konfigurasi_transmisi" form:"konfigurasi_transmisi"`
}
