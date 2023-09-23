package respon

import (
	"17/eksplorasi/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// Rute-rute untuk books
	e.GET("/books", controller.GetbookController)
	e.GET("/books/:id", controller.GetBookController)
	e.POST("/books", controller.CreateBookController)
	e.DELETE("/books/:id", controller.DeleteBookController)
	e.PUT("/books/:id", controller.UpdateBookController)

	// Rute-rute untuk users
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)

	// Rute-rute untuk blogs
	e.GET("/blogs", controller.GetAllBlogController)
	e.GET("/blogs/:id", controller.GetBlogController)
	e.POST("/blogs", controller.CreateBlogsController)
	e.DELETE("/blogs/:id", controller.DeleteBlogController)
	e.PUT("/blogs/:id", controller.UpdateBlogController)

	return e
}
