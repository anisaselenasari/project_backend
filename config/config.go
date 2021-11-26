package config

import (
	"project_backend/db"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	//declare struct config & variable connectionString
	dsn := "selena:admin@tcp(host.docker.internal:3306)/review_mobil?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	initMigration()
}

var (
	DB *gorm.DB
)

func initMigration() {
	DB.AutoMigrate(
		&db.Admins{},
		&db.Descs{},
		&db.Dimensi{},
		&db.Eksterior{},
		&db.Hiburan{},
		&db.Kenyamanan{},
		&db.Keselamatan{},
		&db.Merks{},
		&db.Mesin{},
		&db.Performs{},
		&db.Types{},
	)
}
