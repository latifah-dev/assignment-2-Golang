package main

import (
	"Tugas2/database"
	"Tugas2/routers"
)

func main() {
	database.StartDB()
	routers.RootHandler().Run(":8080")
}
