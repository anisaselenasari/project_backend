package controller

//import packages
import (
	"net/http"
	"project_backend/config"
	"project_backend/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data deskripsi smartphone
func GetDescs(c echo.Context) error {
	var desc []db.Descs

	if err := config.DB.Find(&desc).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Deskripsi Mobil",
		"descs":   desc,
	})
}

//Fungsi get deskripsi by ID
func GetDescsByID(c echo.Context) error {
	var desc db.Descs
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&desc).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Deskripsi Mobil ",
		"descs":   desc,
	})
}

//fungsi create new deskripsi
func CreateDescs(e echo.Context) error {
	desc := db.Descs{}
	e.Bind(&desc)

	if err := config.DB.Save(&desc).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menambahkan data deskripsi mobil",
		"descs":   desc,
	})
}

//Fungsi Update Tabel Deskripsi
func UpdateDescsByID(e echo.Context) error {
	desc := db.Descs{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&desc)
	desc.ID = id

	if err := config.DB.Where("id= ?", desc.ID).Updates(&desc).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Deskripsi mobil",
		"descs":   desc,
	})
}

//Fungsi hapus data deskripsi
func DeleteDescsByID(e echo.Context) error {
	var desc db.Descs
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&desc)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"descs":   desc,
		"message": "Data Berhasil Dihapus",
	})
}
