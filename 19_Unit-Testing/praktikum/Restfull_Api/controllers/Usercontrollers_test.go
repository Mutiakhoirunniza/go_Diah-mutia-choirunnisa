package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"unittestting/Restfull_Api/config"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}


func generateAuthToken() (string, error) {
	loginRequest := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    "mutiakhoirunniza@gmail.com",
		Password: "mutiakhoirunniza2002",
	}

	requestBody, err := json.Marshal(loginRequest)
	if err != nil {
		return "", err
	}

	req := httptest.NewRequest(http.MethodPost, "users/", bytes.NewBuffer(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := LoginUserController(c); err != nil {
		return "", err
	}

	if rec.Code != http.StatusBadRequest {
		return "", fmt.Errorf("Login failed with status code %d", rec.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		return "", err
	}

	token, ok := response["token"].(string)
	if !ok {
		return "", errors.New("Token not found in response")
	}

	return token, nil
}

func TestCreateUserController(t *testing.T) {
	e := InitEchoTestAPI()

	reqBody := `{
		"name": "mutia khoirunniza",
		"email": "mutiakhoirunniza@gmail.com",
		"password": "mutiakhoirunniza2002"
	}`
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

// func TestLoginUserController(t *testing.T) {
// 	e := initEchoTestAPI()
// 	reqBody := `{
// 		"email": "mutiakhoirunniza@gmail.com",
// 		"password": "mutiakhoirunniza2002"
// 	}`
// 	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(reqBody))
// 	req.Header.Set("Content-Type", "application/json")
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	if assert.NoError(t, LoginUserController(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 	}
// }

func TestGetUsersAllController(t *testing.T) {
	e := echo.New()

	req1 := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)

	err1 := GetUsersAllController(c1)

	if assert.NoError(t, err1) {
		assert.Equal(t, http.StatusOK, rec1.Code)
	} else {
		fmt.Println(rec1.Body.String())
	}

	req2 := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)

	err2 := GetUsersAllController(c2)

	if assert.NoError(t, err2) {
		assert.Equal(t, http.StatusOK, rec2.Code)
	}

	req3 := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)

	err3 := GetUsersAllController(c3)

	if assert.NoError(t, err3) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}
}

func TestGetUserController(t *testing.T) {
	e := echo.New()
	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

func TestUpdateUserController(t *testing.T) {
	e := echo.New()
	t.Run("InvalidToken", func(t *testing.T) {
		reqBody := `{
			"name": "arnita argianata",
			"email": "nita@email.com",
			"password": "2009"
		}`

		req := httptest.NewRequest(http.MethodPut, "/users/:id", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("NoToken", func(t *testing.T) {
		reqBody := `{
			"name": "arnita argianata",
			"email": "nita@email.com",
			"password": "2009"
		}`

		req := httptest.NewRequest(http.MethodPut, "/users/:id", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

func TestDeleteUserController(t *testing.T) {
	e := echo.New()
	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("7")

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
