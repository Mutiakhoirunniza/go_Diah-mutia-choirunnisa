package router

import (
	"17/eksplorasi/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	e.GET("/books", controller.GetbookController)
	e.GET("/books/:id", controller.GetBookController)
	e.POST("/books", controller.CreateBookController)
	e.DELETE("/books/:id", controller.DeleteBookController)
	e.PUT("/books/:id", controller.UpdateBookController)

	return e
}

func new() *echo.Echo {

	e := echo.New()

	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)

	return e
}
