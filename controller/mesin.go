package controller

//import packages
import (
	"net/http"
	"project_backend/config"
	"project_backend/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data mesin
func GetMesin(c echo.Context) error {
	var mesin []db.Mesin

	if err := config.DB.Find(&mesin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Mesin Mobil",
		"mesin":   mesin,
	})
}

//Fungsi get mesin by ID
func GetMesinByID(c echo.Context) error {
	var mesin db.Mesin
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&mesin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Mesin Mobil ",
		"mesin":   mesin,
	})
}

//fungsi create new mesin
func CreateMesin(e echo.Context) error {
	mesin := db.Mesin{}
	e.Bind(&mesin)

	if err := config.DB.Save(&mesin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menambahkan data mesin mobil",
		"mesin":   mesin,
	})
}

//Fungsi Update Tabel mesin
func UpdateMesinByID(e echo.Context) error {
	mesin := db.Mesin{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&mesin)
	mesin.ID = id

	if err := config.DB.Where("id= ?", mesin.ID).Updates(&mesin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data mesin mobil",
		"mesin":   mesin,
	})
}

//Fungsi hapus data mesin
func DeleteMesinByID(e echo.Context) error {
	var mesin db.Mesin
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&mesin)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"mesin":   mesin,
		"message": "Data Berhasil Dihapus",
	})
}
