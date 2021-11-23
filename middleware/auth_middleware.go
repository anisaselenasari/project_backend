// package middleware

// import (
// 	"project_backend/config"
// 	"project_backend/db"

// 	"github.com/labstack/echo/v4"
// )

// func AuthDB(adminname, password string, c echo.Context) (bool, error) {
// 	var admin db.Admins

// 	err := config.DB.Where("email = ? AND password = ?", adminname, password).First(&admin).Error
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
