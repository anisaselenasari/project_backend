package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	//declare struct config & variable connectionString
	dsn := "sele:admin@tcp(127.0.0.1)/review_mobil?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 	initMigration()
}

var (
	DB *gorm.DB
)

// func initMigration() {
// 	DB.AutoMigrate(&db.Admins{}
// 		)
// }
