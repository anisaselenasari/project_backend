package controller

//import packages
import (
	"net/http"
	"project_backend/config"
	"project_backend/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data dimensi
func GetDimensi(c echo.Context) error {
	var dimensi []db.Dimensi

	if err := config.DB.Find(&dimensi).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Dimensi Mobil",
		"dimensi": dimensi,
	})
}

//Fungsi get dimensiby ID
func GetDimensiByID(c echo.Context) error {
	var dimensi db.Dimensi
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&dimensi).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Dimensi Mobil ",
		"dimensi": dimensi,
	})
}

//fungsi create new dimensi
func CreateDimensi(e echo.Context) error {
	dimensi := db.Dimensi{}
	e.Bind(&dimensi)

	if err := config.DB.Save(&dimensi).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menambahkan data dimensi",
		"dimensi": dimensi,
	})
}

//Fungsi Update Tabel Dimensi
func UpdateDimensiByID(e echo.Context) error {
	dimensi := db.Dimensi{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&dimensi)
	dimensi.ID = id

	if err := config.DB.Where("id= ?", dimensi.ID).Updates(&dimensi).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Dimensi",
		"dimensi": dimensi,
	})
}

//Fungsi hapus data dimensi
func DeleteDimensiByID(e echo.Context) error {
	var dimensi db.Dimensi
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&dimensi)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"dimensi": dimensi,
		"message": "Data Berhasil Dihapus",
	})
}
