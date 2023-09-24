package controllers

import (
	"net/http"
	"strconv"
	"widellware/config"
	"widellware/models"

	"github.com/labstack/echo/v4"
)

// get all Book
func GetbookAllController(c echo.Context) error {
	var book []models.Book

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

	var book models.Book
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
	Book := models.Book{}
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

	var book models.Book
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

	book := new(models.Book)
	if err := c.Bind(book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var existingBook models.Book
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
