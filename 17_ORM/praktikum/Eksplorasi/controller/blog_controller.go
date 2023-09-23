package controller

import (
	"17/Eksplorasi/config"
	"17/Eksplorasi/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateBlogsController(c echo.Context) error {
	blog := new(model.Blog)
	if err := c.Bind(blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user model.User
	if err := config.DB.First(&user, blog.UserID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid UserID")
	}

	if err := config.DB.Create(blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"blog":    blog,
	})
}

func GetAllBlogController(c echo.Context) error {
	var blogs []model.Blog
	if err := config.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// GetBlogController melihat rincian data blog berdasarkan ID
func GetBlogController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var blog model.Blog
	if err := config.DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog",
		"blog":    blog,
	})
}

func UpdateBlogController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	blog := new(model.Blog)
	if err := c.Bind(blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var existingBlog model.Blog
	if err := config.DB.First(&existingBlog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}

	existingBlog.Title = blog.Title
	existingBlog.Content = blog.Content

	if err := config.DB.Save(&existingBlog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog",
		"blog":    existingBlog,
	})
}

func DeleteBlogController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	var blog model.Blog
	if err := config.DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}

	if err := config.DB.Delete(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog",
	})
}
