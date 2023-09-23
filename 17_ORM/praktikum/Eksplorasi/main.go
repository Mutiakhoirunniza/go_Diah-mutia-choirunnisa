package main

import (
	"17/Eksplorasi/config"
	"17/Eksplorasi/router"
)

func main() {
	// create a new echo instance
	config.InitDB()

	e := router.New()
	e.Start(":8000")
}
