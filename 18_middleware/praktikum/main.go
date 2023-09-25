package main

import (
	"widellware/config"
	"widellware/routes"
)

func main() {
	// create a new echo instance
	config.InitDB()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
