package main

import (
	"project_backend/config"
)

//MAIN FUNCTION
func main() {
	config.InitDB()
	e.Start(":8080")
}
