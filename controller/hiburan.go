package controller

//import packages
import (
	"net/http"
	"project_backend/config"
	"project_backend/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data hiburan
func GetHiburan(c echo.Context) error {
	var hiburan []db.Hiburan

	if err := config.DB.Find(&hiburan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Hiburan Mobil",
		"hiburan": hiburan,
	})
}

//Fungsi get hiburan by ID
func GetHiburanByID(c echo.Context) error {
	var hiburan db.Hiburan
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&hiburan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Hiburan Mobil ",
		"hiburan": hiburan,
	})
}

//fungsi create new hiburan
func CreateHiburan(e echo.Context) error {
	hiburan := db.Hiburan{}
	e.Bind(&hiburan)

	if err := config.DB.Save(&hiburan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Menambahkan Data Hiburan",
		"hiburan": hiburan,
	})
}

//Fungsi Update Tabel Hiburan
func UpdateHiburanByID(e echo.Context) error {
	hiburan := db.Hiburan{}
	e.Bind(&hiburan)

	if err := config.DB.Where("id= ?", hiburan.ID).Updates(&hiburan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Hiburan",
		"hiburan": hiburan,
	})
}

//Fungsi hapus data hiburan
func DeleteHiburanByID(e echo.Context) error {
	var hiburan db.Hiburan
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&hiburan)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"hiburan": hiburan,
		"message": "Data Berhasil Dihapus",
	})
}
