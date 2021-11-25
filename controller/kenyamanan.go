package controller

//import packages
import (
	"net/http"
	"project_backend/config"
	"project_backend/db"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data kenyamanan
func GetKenyamanan(c echo.Context) error {
	var kenyamanan []db.Kenyamanan

	if err := config.DB.Find(&kenyamanan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Berhasil Menampilkan Semua Data Kenyamanan Mobil",
		"kenyamanan": kenyamanan,
	})
}

//Fungsi get kenyamanan by ID
func GetKenyamananByID(c echo.Context) error {
	var kenyamanan db.Kenyamanan
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := config.DB.Where("id= ?", id).Find(&kenyamanan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Berhasil Menampilkan Kenyamanan Mobil ",
		"kenyamanan": kenyamanan,
	})
}

//fungsi create new kenyamanan
func CreateKenyamanan(e echo.Context) error {
	kenyamanan := db.Kenyamanan{}
	e.Bind(&kenyamanan)

	if err := config.DB.Save(&kenyamanan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Success Menambahkan Data Kenyamanan mobil",
		"kenyamanan": kenyamanan,
	})
}

//Fungsi Update Tabel Kenyamanan
func UpdatKenyamananByID(e echo.Context) error {
	kenyamanan := db.Kenyamanan{}
	e.Bind(&kenyamanan)

	if err := config.DB.Updates(&kenyamanan).Where("id= ?", kenyamanan.ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Berhasil Mengubah Data Kenyamanan Mobil",
		"kenyamanan": kenyamanan,
	})
}

//Fungsi hapus data kenyamanan
func DeleteKenyamananByID(e echo.Context) error {
	var kenyamanan db.Kenyamanan
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&kenyamanan)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"kenyamanan": kenyamanan,
		"message":    "Data Berhasil Dihapus",
	})
}
