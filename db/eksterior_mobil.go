package db

type Eksterior struct {
	ID               int    `json:"id" form:"id"`
	Lampu_KabutDepan string `json:"lampu_kabutdepan" form:"lampu_kabutdepan"`
	Lampu_Depan      string `json:"lampu_depan" form:"lampu_depan"`
	Kaca_Spion       string `json:"kaca_spion" form:"kaca_spion"`
	Lampu_Belakang   string `json:"lampu_belakang" form:"lampu_belakang"`
}
