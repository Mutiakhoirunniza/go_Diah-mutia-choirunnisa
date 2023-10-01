package routes

import (
	"unittestting/Restfull_Api/controllers"
	m "unittestting/Restfull_Api/middleware"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	protectUsers := e.Group("/users", m.JWTMiddleware())

	protectUsers.GET("", controllers.GetUsersAllController)
	protectUsers.GET("/:id", controllers.GetUserController)
	protectUsers.DELETE("/:id", controllers.DeleteUserController)
	protectUsers.PUT("/:id", controllers.UpdateUserController)
	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)

	protectBooks := e.Group("/books", m.JWTMiddleware())

	protectBooks.GET("", controllers.GetbookAllController)
	protectBooks.GET("/:id", controllers.GetBookController)
	protectBooks.PUT("/:id", controllers.UpdateBookController)
	protectBooks.DELETE("/:id", controllers.DeleteBookController)
	protectBooks.POST("", controllers.CreateBookController)

	return e
}
