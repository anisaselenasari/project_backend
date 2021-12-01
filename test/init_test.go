package test

import (
	"fmt"
	"net/http/httptest"
	"project_backend/config"
	"project_backend/db"
	"project_backend/route"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/labstack/echo/v4"
)

var (
	echoHandler *echo.Echo
	server      *httptest.Server
)

func init() {
	config.InitDB("project_backend_test")

	echoHandler = route.New()
	server = httptest.NewServer(echoHandler)

}

func TearUp() func() {
	config.InitDB("project_backend_test")

	article := &db.Types{
		Merk: "Test1",
		Tipe: "Test2",
	}
	if err := config.DB.Create(&article).Error; err != nil {
		fmt.Println("error")
	}
	article1 := &db.Merks{
		Merk: "Test3",
	}
	if err := config.DB.Create(&article1).Error; err != nil {
		fmt.Println("error")
	}
	return func() {
		config.InitDB("project_backend_test")
		config.DB.Exec("TRUNCATE TABLE types;")
		config.DB.Exec("TRUNCATE TABLE merks;")
	}
}

// GetHttpExpect Get cached expect, create new if nil
func GetHTTPExpect(t *testing.T) *httpexpect.Expect {
	if server == nil {
		server = httptest.NewServer(echoHandler)
	}
	return NewHTTPExpect(t)
}

func NewHTTPExpect(t *testing.T) *httpexpect.Expect {
	// TODO : printer set, only if the testing is failed
	// https://github.com/gavv/httpexpect/issues/69
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: server.URL,
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		},
		Reporter: t,
	})
}
