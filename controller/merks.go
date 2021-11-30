package controller

//import packages
import (
	"net/http"
	"project_backend/config"
	"project_backend/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data admin
func GetMerks(c echo.Context) error {
	var merks []db.Merks

	if err := config.DB.Find(&merks).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Merk Mobil",
		"merks":   merks,
	})
}

//Fungsi get admin by ID
func GetMerksByID(c echo.Context) error {
	var merks db.Merks
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&merks).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Merk Mobil ",
		"merks":   merks,
	})
}

//fungsi create new admins
func CreateMerk(e echo.Context) error {
	merks := db.Merks{}
	e.Bind(&merks)

	if err := config.DB.Save(&merks).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menambahkan merks  mobil",
		"merks":   merks,
	})
}

//Fungsi Update Tabel Merk HP
func UpdateMerkByID(e echo.Context) error {
	merks := db.Merks{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&merks)
	merks.ID = id

	if err := config.DB.Where("id= ?", merks.ID).Updates(&merks).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data merks",
		"merks":   merks,
	})
}

//Fungsi hapus data merk
func DeleteMerkByID(e echo.Context) error {
	var merks db.Merks
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&merks)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"merks":   merks,
		"message": "Data Berhasil Dihapus",
	})
}
