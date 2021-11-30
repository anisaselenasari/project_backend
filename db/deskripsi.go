package db

type Descs struct {
	ID             int `json:"id" form:"id"`
	Id_Types       int `json:"id_types" form:"id_types"`
	Id_Kenyamanan  int `json:"id_kenyamanan" form:"id_kenyamanan"`
	Id_Keselamatan int `json:"id_keselamatan" form:"id_keselamatan"`
	Id_Eksterior   int `json:"id_eksterior" form:"id_eksterior"`
	Id_Mesin       int `json:"id_mesin" form:"id_mesin"`
	Id_Performa    int `json:"id_performa" form:"id_performa"`
	Id_Dimensi     int `json:"id_dimensi" form:"id_dimensi"`
	Id_Hiburan     int `json:"id_hiburan" form:"id_hiburan"`
}
