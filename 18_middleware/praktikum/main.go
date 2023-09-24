package main

import (
	"widellware/config"
	"widellware/routes"
)

func main() {
	// create a new echo instance
	config.InitDB()

	e := routes.New()
	e.Start(":8000")
}
