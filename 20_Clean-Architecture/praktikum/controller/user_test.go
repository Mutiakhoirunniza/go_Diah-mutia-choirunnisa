package controller_test

import (
	"20/controller"
	"20/model"
	"20/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

)

func TestGetAllUsers(t *testing.T) {
	e := echo.New()
	mockUserRepo := &repository.MockUserRepository{}
	uController := controller.NewUserController(mockUserRepo)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUserRepo.On("Find").Return([]model.User{}, nil).Once()

	if assert.NoError(t, uController.GetAllUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	mockUserRepo.AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	reqbody := `{
		"email": "mutiakhoirunniza@gmail.com",
		"password":"2002"}`

	mockUserRepo := &repository.MockUserRepository{}
	uController := controller.NewUserController(mockUserRepo)
	mockUserRepo.On("Create", mock.AnythingOfType("model.User")).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqbody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	err := uController.CreateUser(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUserRepo.AssertExpectations(t)
}

func TestCreateUserfailed(t *testing.T) {

	mockUserRepo := &repository.MockUserRepository{}
	uController := controller.NewUserController(mockUserRepo)

	mockUserRepo.On("Create", mock.AnythingOfType("model.User")).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	err := uController.CreateUser(c)

	assert.NoError(t, err)

	mockUserRepo.AssertExpectations(t)
}


// // someValidationError adalah tipe error kustom
// type someValidationError struct {
//     // Anda bisa menambahkan field tambahan jika diperlukan
//     Message string
// }

// // Implementasi method Error() untuk tipe someValidationError.
// func (e *someValidationError) Error() string {
//     return e.Message
// }

// // someUnauthorizedError adalah tipe error kustom
// type someUnauthorizedError struct {
//     // Anda bisa menambahkan field tambahan jika diperlukan
//     Message string
// }

// // Implementasi method Error() untuk tipe someUnauthorizedError.
// func (e *someUnauthorizedError) Error() string {
//     return e.Message
// }





// func TestCreateUser_WithValidationError(t *testing.T) {
// 	reqbody := `{
//         "email": "invalid-email",
//         "password":"2002"
//     }`

// 	mockUserRepo := &repository.MockUserRepository{}
// 	uController := controller.NewUserController(mockUserRepo)

// 	mockUserRepo.On("Create", mock.AnythingOfType("model.User")).Return(e.Message)

// 	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqbody))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	e := echo.New()
// 	c := e.NewContext(req, rec)

// 	err := uController.CreateUser(c)

// 	assert.Error(t, err)
// 	assert.Equal(t, http.StatusBadRequest, rec.Code)

// 	mockUserRepo.AssertExpectations(t)
// }

// func TestCreateUser_WithUnauthorizedError(t *testing.T) {
// 	reqbody := `{
//         "email": "mutiakhoirunniza@gmail.com",
//         "password":"2002"
//     }`

// 	mockUserRepo := &repository.MockUserRepository{}
// 	uController := controller.NewUserController(mockUserRepo)

// 	// Set up the mockUserRepo to return an unauthorized error.

// 	mockUserRepo.On("Create", mock.AnythingOfType("model.User")).Return(someUnauthorizedError)

// 	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqbody))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	e := echo.New()
// 	c := e.NewContext(req, rec)

// 	err := uController.CreateUser(c)

// 	assert.Error(t, err)
// 	assert.Equal(t, http.StatusUnauthorized, rec.Code)

// 	mockUserRepo.AssertExpectations(t)
// }
