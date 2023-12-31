package controllers

import (
	"net/http"
	"strconv"
	"unittestting/Restfull_Api/config"
	"unittestting/Restfull_Api/middleware"
	"unittestting/Restfull_Api/models"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersAllController(c echo.Context) error {
	var users []models.Users

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

	var user models.Users
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
	user := models.Users{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func LoginUserController(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)

	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Fail login",
			"error":   err.Error(),
		})
	}

	userID := int(user.ID)

	token, err := middleware.CreateToken(userID, user.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Fail login",
			"error":   err.Error(),
		})
	}

	UsersResponse := models.UsersResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    UsersResponse,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	IdStr := c.Param("id")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	var user models.Users
	// if err := config.DB.First(&user, Id).Error; err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "User Not Found")
	// }

	if err := config.DB.Delete(&user,Id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
		// "user":    user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	IdStr := c.Param("id")
	Id, err := strconv.Atoi(IdStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	user := new(models.Users)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var existingUser models.Users
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

