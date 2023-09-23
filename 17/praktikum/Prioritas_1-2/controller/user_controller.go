package controller

import (
	"17/Prioritas_1-2/config"
	"17/Prioritas_1-2/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	IdStr := c.Param("id")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	var user model.User
	if err := config.DB.First(&user, Id).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, "User Not Found")

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by ID",
		"user":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	IdStr := c.Param("id")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	var user model.User
	if err := config.DB.First(&user, Id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User Not Found")
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
		"user":    user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	IdStr := c.Param("id")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var existingUser model.User
	if err := config.DB.First(&existingUser, Id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password

	if err := config.DB.Save(&existingUser).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    existingUser,
	})
}
