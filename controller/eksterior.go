package controller

//import packages
import (
	"net/http"
	"project_backend/config"
	"project_backend/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data eksterior
func GetEksterior(c echo.Context) error {
	var eksterior []db.Eksterior

	if err := config.DB.Find(&eksterior).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Berhasil Menampilkan Semua Data Eksterior Mobil",
		"eksterior": eksterior,
	})
}

//Fungsi get eksterior by ID
func GetEksteriorByID(c echo.Context) error {
	var eksterior db.Eksterior
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&eksterior).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Berhasil Menampilkan Eksterior Mobil ",
		"eksterior": eksterior,
	})
}

//fungsi create new eksterior
func CreateEksterior(e echo.Context) error {
	eksterior := db.Eksterior{}
	e.Bind(&eksterior)

	if err := config.DB.Save(&eksterior).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Success Menambahkan Data Eksterior",
		"eksterior": eksterior,
	})
}

//Fungsi Update Tabel Eksterior
func UpdateEksteriorByID(e echo.Context) error {
	eksterior := db.Eksterior{}
	e.Bind(&eksterior)

	if err := config.DB.Updates(&eksterior).Where("id= ?", eksterior.ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Berhasil Mengubah Data Eksterior",
		"eksterior": eksterior,
	})
}

//Fungsi hapus data eksterior
func DeletePerfomByID(e echo.Context) error {
	var perfoms db.Performs
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&perfoms)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"perfoms": perfoms,
		"message": "Data Berhasil Dihapus",
	})
}
