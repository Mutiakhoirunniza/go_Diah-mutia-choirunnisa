package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// routing (request)
	e.GET("/coba/:id/:ege", Usercontroller)
	e.Start(":8000")
}

type User struct {
	Id 	  int
	Age   int
	Email string
	Name  string
}

// controller (respon)
func Usercontroller(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	age, _ := strconv.Atoi(e.Param("ege"))

	user := User{Id: id, Age: age, Email: "mutiakhoirunniza@gmail.com", Name: "mutea"}

	return e.JSON(http.StatusOK, user)
}
