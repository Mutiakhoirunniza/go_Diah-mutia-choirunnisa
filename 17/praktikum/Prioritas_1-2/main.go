package main

import (
	"17/Prioritas_1-2/config"
	"17/Prioritas_1-2/router"
)

func main() {
	// create a new echo instance
	config.InitDB()

	e := router.New()
	e.Start(":8000")
}
