package main

import (
	"fmt"
	"laptop/handler"
	"laptop/usecase"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load .env file. Err: %s", err)
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	laptopUsecase := usecase.NewLaptopUsecase()
	laptopHandler := handler.NewLaptopHandler(laptopUsecase)
	e.POST("/recommendation-Laptop", laptopHandler.RecommendLaptop)
	port := ":8800"
	fmt.Printf("Server is running on port %s\n", port)
	e.Start(port)
}
