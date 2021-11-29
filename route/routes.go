package route

import (
	"project_backend/constants"
	"project_backend/controller"
	"project_backend/middleware"
	m "project_backend/middleware"

	"github.com/labstack/echo/v4"
	midd "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	//CRUD ADMIN
	e.GET("/admins", controller.GetAdminsController)
	e.GET("/admins/:id", controller.GetAdminsByID)
	m.LogMiddleware(e)
	e.POST("/admins", controller.CreateAdminController)
	e.DELETE("/admins/:id", controller.DeleteAdminByID)
	e.PUT("/admins/:id", controller.UpdateAdminByID)
	e.POST("/login", controller.LoginAdminController)

	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(midd.BasicAuth(middleware.BasicAuthDB))
	eAuthBasic.GET("/admins", controller.GetAdminsController)

	eJwt := e.Group("/jwt")
	eJwt.Use(midd.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/admins", controller.GetAdminsController)

	//CRUD MERK
	e.GET("/merks", controller.GetMerks)
	e.GET("/merks/:id", controller.GetMerksByID)
	e.POST("/merks", controller.CreateMerk)
	e.DELETE("/merks/:id", controller.DeleteMerkByID)
	e.PUT("/merks/:id", controller.UpdateMerkByID)

	//CRUD Type Mobil
	e.GET("/types", controller.GetTypes)
	e.GET("/types/:id", controller.GetTypesByID)
	e.POST("/types", controller.CreateType)
	e.DELETE("/types/:id", controller.DeleteTypeByID)
	e.PUT("/types/:id", controller.UpdateTypeByID)

	// //CRUD Advances Smartphone
	// e.GET("/advances", controller.GetAdvances)
	// e.GET("/advances/:id", controller.GetAdvancesByID)
	// e.POST("/advances", controller.CreateAdvances)
	// e.DELETE("/advances/:id", controller.DeleteAdvanceByID)
	// e.PUT("/advances/:id", controller.UpdateAdvanceByID)

	//CRUD Descs Mobil
	e.GET("/descs", controller.GetDescs)
	e.GET("/descs/:id", controller.GetDescsByID)
	e.POST("/descs", controller.CreateDescs)
	e.DELETE("/descs/:id", controller.DeleteDescsByID)
	e.PUT("/descs/:id", controller.UpdateDescsByID)

	//CRUD Dimensi Mobil
	e.GET("/dimensi", controller.GetDimensi)
	e.GET("/dimensi/:id", controller.GetDimensiByID)
	e.POST("/dimensi", controller.CreateDimensi)
	e.DELETE("/dimensi/:id", controller.DeleteDimensiByID)
	e.PUT("dimensi/:id", controller.UpdateDimensiByID)

	//CRUD Eksterior Mobil
	e.GET("/eksterior", controller.GetEksterior)
	e.GET("/eksterior/:id", controller.GetEksteriorByID)
	e.POST("/eksterior", controller.CreateEksterior)
	e.DELETE("/eksterior/:id", controller.DeleteEksteriorByID)
	e.PUT("/eksterior/:id", controller.UpdateEksteriorByID)

	//CRUD Hiburan Mobil
	e.GET("/hiburan", controller.GetHiburan)
	e.GET("/hiburan/:id", controller.GetHiburanByID)
	e.POST("/hiburan", controller.CreateHiburan)
	e.DELETE("/hiburan/:id", controller.DeleteHiburanByID)
	e.PUT("/hiburan/:id", controller.UpdateHiburanByID)

	//CRUD Kenyamanan Mobil
	e.GET("/kenyamanan", controller.GetKenyamanan)
	e.GET("/kenyamanan/:id", controller.GetKenyamananByID)
	e.POST("/kenyamanan", controller.CreateKenyamanan)
	e.DELETE("/kenyamanan/:id", controller.DeleteKenyamananByID)
	e.PUT("/kenyamanan/:id", controller.UpdatKenyamananByID)

	//CRUD Keselamatan Mobil
	e.GET("/keselamatan", controller.GetKeselamatan)
	e.GET("/keselamatan/:id", controller.GetKeselamatanByID)
	e.POST("/keselamatan", controller.CreateKeselamatan)
	e.DELETE("/keselamatan/:id", controller.DeleteKeselamatanByID)
	e.PUT("/keselamatan/:id", controller.UpdateKeselamatanByID)

	//CRUD Performa Mobil
	e.GET("/performs", controller.GetPerfoms)
	e.GET("/performs/:id", controller.GetPerfomsByID)
	e.POST("/performs", controller.CreatePerfoms)
	e.DELETE("/performs/:id", controller.DeletePerfomsByID)
	e.PUT("/performs/:id", controller.UpdatePerfomsByID)

	//CRUD Mesin Mobil
	e.GET("/mesin", controller.GetMesin)
	e.GET("/mesin/:id", controller.GetMesinByID)
	e.POST("/mesin", controller.CreateMesin)
	e.DELETE("/mesin/:id", controller.DeleteMesinByID)
	e.PUT("/mesin/:id", controller.UpdateMesinByID)

	return e
}
