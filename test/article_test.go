package test

import (
	"net/http"
	"testing"
)

func TestGetType(t *testing.T) {
	tearDown := TearUp()
	defer tearDown()

	e := GetHTTPExpect(t)

	result := e.GET("/types").Expect().Status(http.StatusOK).JSON().Object()

	result.Value("message").String().Contains("Berhasil Menampilkan Semua Tipe Mobil")
	result.Value("types").Array().NotEmpty()
}

func TestGetMerk(t *testing.T) {
	tearDown := TearUp()
	defer tearDown()

	e := GetHTTPExpect(t)

	result := e.GET("/merks").Expect().Status(http.StatusOK).JSON().Object()

	result.Value("message").String().Contains("Berhasil Menampilkan Semua Merk Mobil")
	result.Value("merks").Array().NotEmpty()
}
