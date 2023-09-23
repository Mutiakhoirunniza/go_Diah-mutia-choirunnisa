package controller

import (
	"17/PRIORITAS_2-1/config"
	"17/PRIORITAS_2-1/model"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all Book
func GetbookController(c echo.Context) error {
	var book []model.Book

	if err := config.DB.Find(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all Book",
		"book":    book,
	})
}

// get Book by id
func GetBookController(c echo.Context) error {
	IdStr := c.Param("id")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Book ID")
	}

	var book model.Book
	if err := config.DB.First(&book, Id).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, "book Not Found")

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by ID",
		"book":    book,
	})
}

// create new Book
func CreateBookController(c echo.Context) error {
	Book := model.Book{}
	c.Bind(&Book)

	if err := config.DB.Save(&Book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    Book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	IdStr := c.Param("id")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid book ID")
	}

	var book model.Book
	if err := config.DB.First(&book, Id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "book Not Found")
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
		"book":    book,
	})
}

// update Book by id
func UpdateBookController(c echo.Context) error {
	IdStr := c.Param("id")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	book := new(model.Book)
	if err := c.Bind(book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var existingBook model.Book
	if err := config.DB.First(&existingBook, Id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}
	existingBook.Title = book.Title
	existingBook.Author = book.Author
	existingBook.Publisher = book.Publisher

	if err := config.DB.Save(&existingBook).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"user":    existingBook,
	})
}
