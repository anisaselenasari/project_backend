package db

type Dimensi struct {
	ID           int `json:"id" form:"id"`
	Jarak        int `json:"jarak" form:"jarak"`
	TinggiMobil  int `json:"tinggimobil" form:"tinggimobil"`
	PanjangMobil int `json:"panjangmobil" form:"panjangmobil"`
}
