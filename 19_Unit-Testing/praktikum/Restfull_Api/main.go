package main

import (
	"unittestting/Restfull_Api/config"
	"unittestting/Restfull_Api/routes"
)

func main() {
	// create a new echo instance
	config.InitDBTest()

	e := routes.New()
	e.Logger.Fatal(e.Start(":8100"))
}
