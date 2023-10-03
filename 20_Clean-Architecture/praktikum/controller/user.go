package controller

import (
	"20/dto"
	"20/middleware"
	"20/model"
	"20/repository"

	"github.com/labstack/echo/v4"
)

type UserController interface{}

type userController struct {
	userRepo repository.UserRepository
}

func NewUserController(uRepo repository.UserRepository) *userController {
	return &userController{
		uRepo,
	}
}

func (u *userController) GetAllUsers(c echo.Context) error {
	users, err := u.userRepo.Find()
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": users,
	})
}

func (u *userController) CreateUser(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	err = u.userRepo.Create(user)
	if err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	token, err := middleware.CreateToken(user.Email)
	if err != nil {
		return c.JSON(401, echo.Map{
			"error": err.Error(),
		})
	}

	uRes := dto.DTOUsers{
		Email: user.Email,
		Token: token,
	}
	return c.JSON(200, echo.Map{
		"data": uRes,
	})
}
