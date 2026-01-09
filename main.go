package main

import (
	"laporpak/database"
	"laporpak/router"
)

func main() {
	database.Connect()
	router.StartServer()

}
