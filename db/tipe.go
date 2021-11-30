package db

type Types struct {
	ID       int    `json:"id" form:"id"`
	Id_Merk  int    `json:"id_merks" form:"id_merk"`
	Merk     string `json:"merk"`
	Tipe     string `json:"tipe" form:"tipe"`
	Th_Rilis string `json:"th_rilis" form:"th_rilis"`
	Harga    string `json:"harga" form:"harga"`
}
