package controller

//import packages
import (
	"net/http"
	"project_backend/config"
	"project_backend/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data perfoms
func GetPerfoms(c echo.Context) error {
	var perfoms []db.Performs

	if err := config.DB.Find(&perfoms).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Perfoms Mobil",
		"perfoms": perfoms,
	})
}

//Fungsi get perfoms by ID
func GetPerfomsByID(c echo.Context) error {
	var perfoms db.Performs
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&perfoms).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Perfoms Mobil ",
		"perfoms": perfoms,
	})
}

//fungsi create new perfoms
func CreatePerfoms(e echo.Context) error {
	perfoms := db.Performs{}
	e.Bind(&perfoms)

	if err := config.DB.Save(&perfoms).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menambahkan data perfoms mobil",
		"perfoms": perfoms,
	})
}

//Fungsi Update Tabel Perfoms
func UpdatePerfomsByID(e echo.Context) error {
	perfoms := db.Performs{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&perfoms)
	perfoms.ID = id

	if err := config.DB.Where("id= ?", perfoms.ID).Updates(&perfoms).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data perfoms mobil",
		"perfoms": perfoms,
	})
}

//Fungsi hapus data perfoms
func DeletePerfomsByID(e echo.Context) error {
	var perfoms db.Performs
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&perfoms)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"perfoms": perfoms,
		"message": "Data Berhasil Dihapus",
	})
}
