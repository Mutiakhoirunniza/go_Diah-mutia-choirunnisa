package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Users struct {
	ID       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []Users

func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

func GetUserController(c echo.Context) error {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid user ID",
		})
	}

	var foundUser Users
	for _, user := range users {
		if user.ID == userID {
			foundUser = user
			break
		}
	}

	if foundUser.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}

	return c.JSON(http.StatusOK, foundUser)
}

func DeleteUserController(c echo.Context) error {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid user ID",
		})
	}

	var indexToRemove = -1
	for i, user := range users {
		if user.ID == userID {
			indexToRemove = i
			break
		}
	}

	if indexToRemove == -1 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}

	users = append(users[:indexToRemove], users[indexToRemove+1:]...)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}

func UpdateUserController(c echo.Context) error {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid user ID",
		})
	}

	var indexToUpdate = -1
	for i, user := range users {
		if user.ID == userID {
			indexToUpdate = i
			break
		}
	}

	if indexToUpdate == -1 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}

	var updatedUser Users
	if err := c.Bind(&updatedUser); err != nil {
		return err
	}

	updatedUser.ID = userID
	users[indexToUpdate] = updatedUser

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
		"user":    updatedUser,
	})
}

func CreateUserController(c echo.Context) error {
	user := Users{}
	c.Bind(&user)

	if len(users) == 0 {
		user.ID = 1
	} else {
		newID := users[len(users)-1].ID + 1
		user.ID = newID
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func main() {
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}
