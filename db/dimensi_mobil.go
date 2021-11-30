package db

type Dimensi struct {
	ID           int    `json:"id" form:"id"`
	Jarak        string `json:"jarak" form:"jarak"`
	TinggiMobil  string `json:"tinggimobil" form:"tinggimobil"`
	PanjangMobil string `json:"panjangmobil" form:"panjangmobil"`
}
