package main

import (
	"17/prioritas_2-1/config"
	"17/prioritas_2-1/router"
)

func main() {
	// create a new echo instance
	config.InitDB()

	e := router.New()
	e.Start(":8000")
}
