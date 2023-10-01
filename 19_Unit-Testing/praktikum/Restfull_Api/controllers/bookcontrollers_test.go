package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetbookAllController(t *testing.T) {
	e := echo.New()

	req1 := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec1 := httptest.NewRecorder()
	c1 := e.NewContext(req1, rec1)

	err1 := GetbookAllController(c1)

	if assert.NoError(t, err1) {
		assert.Equal(t, http.StatusOK, rec1.Code)
	} else {
		fmt.Println(rec1.Body.String())
	}

	req2 := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req2, rec2)

	err2 := GetUsersAllController(c2)

	if assert.NoError(t, err2) {
		assert.Equal(t, http.StatusOK, rec2.Code)
	}

	req3 := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec3 := httptest.NewRecorder()
	c3 := e.NewContext(req3, rec3)

	err3 := GetUsersAllController(c3)

	if assert.NoError(t, err3) {
		assert.Equal(t, http.StatusOK, rec3.Code)
	}
}

func TestCreateBookController(t *testing.T) {
	e := InitEchoTestAPI()
	t.Run("InputValid", func(t *testing.T) {
		reqBody := `{
			"title": "The Da Vinci Code",
			"author": "Dan Brown",
			"publisher": "Gramedia"
		}`

		req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		req.Header.Set("Authorization", "InputValid")

		if assert.NoError(t, CreateBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

func TestGetBookController(t *testing.T) {
	e := echo.New()
	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

func TestUpdateBookController(t *testing.T) {
	e := echo.New()
	t.Run("InvalidToken", func(t *testing.T) {
		reqBody := `{
			"title": "The Reckoning",
			"author": " John Grisham",
			"publisher": "Gramedia"
		}`

		req := httptest.NewRequest(http.MethodPut, "/books/:id", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("NoToken", func(t *testing.T) {
		reqBody := `{
			"title": "The Reckoning",
			"author": " John Grisham",
			"publisher": "Gramedia"
		}`

		req := httptest.NewRequest(http.MethodPut, "/books/:id", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

func TestDeleteBookController(t *testing.T) {
	e := echo.New()
	t.Run("InvalidToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/books/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		req.Header.Set("Authorization", "InvalidToken")

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("NoToken", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/books/:id", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
