package main

import (
	"17/eksplorasi/config"
	"17/eksplorasi/respon"
)

func main() {
	// create a new echo instance
	config.InitDB()

	e := respon.New()
	e.Start(":8000")
}
