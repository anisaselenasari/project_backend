package controller

//import packages
import (
	"net/http"
	"project_backend/config"
	"project_backend/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data keselamatan
func GetKeselamatan(c echo.Context) error {
	var keselamatan []db.Keselamatan

	if err := config.DB.Find(&keselamatan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Berhasil Menampilkan Semua Data Keselamatan Mobil",
		"keselamatan": keselamatan,
	})
}

//Fungsi get kenyamanan by ID
func GetKeselamatanByID(c echo.Context) error {
	var keselamatan db.Keselamatan
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&keselamatan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Berhasil Menampilkan Keselamatan Mobil ",
		"keselamatan": keselamatan,
	})
}

//fungsi create new keselamatan
func CreateKeselamatan(e echo.Context) error {
	keselamatan := db.Keselamatan{}
	e.Bind(&keselamatan)

	if err := config.DB.Save(&keselamatan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Success Menambahkan Data Keselamatan mobil",
		"keselamatan": keselamatan,
	})
}

//Fungsi Update Tabel Keselamatan
func UpdateKeselamatanByID(e echo.Context) error {
	keselamatan := db.Keselamatan{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&keselamatan)
	keselamatan.ID = id

	if err := config.DB.Where("id= ?", keselamatan.ID).Updates(&keselamatan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Berhasil Mengubah Data Keselamatan Mobil",
		"keselamatan": keselamatan,
	})
}

//Fungsi hapus data keselamatan
func DeleteKeselamatanByID(e echo.Context) error {
	var keselamatan db.Keselamatan
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&keselamatan)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"keselamatan": keselamatan,
		"message":     "Data Berhasil Dihapus",
	})
}
