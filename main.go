package main

import (
	"project_backend/config"

	"github.com/labstack/echo"
)

//MAIN FUNCTION
func main() {
	config.InitDB()
	e := echo.New()
	e.Start(":8080")
}
