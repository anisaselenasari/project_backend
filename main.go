package main

import (
	"project_backend/config"
	"project_backend/route"
)

//MAIN FUNCTION
func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8181")
}
