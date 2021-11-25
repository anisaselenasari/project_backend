package db

type Hiburan struct {
	ID              int    `json:"id" form:"id"`
	Audio           string `json:"audio" form:"audio"`
	Sistem_Navigasi string `json:"sistem_navigasi" form:"sistem_navigasi"`
}
