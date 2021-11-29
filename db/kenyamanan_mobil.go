package db

type Kenyamanan struct {
	ID                  int    `json:"id" form:"id"`
	AC                  string `json:"ac" form:"ac"`
	Kursi_Lipat         string `json:"kursi_lipat" form:"kursi_lipat"`
	Kapasitas_TmpDuduk  int    `json:"kapasitas_tmpduduk" form:"kapasitas_tmpduduk"`
	Pengaturan_TmpDuduk string `json:"pengaturan_tmpduduk" form:"pengaturan_tmpduduk"`
	Vanity_Mirror       string `json:"vanity_mirror" form:"vanity_mirror"`
}
