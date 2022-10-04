package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	e := echo.New()
	body := `{ "values": ["Sword", "Axe", "Bow"] }`
	req := httptest.NewRequest(http.MethodPost, "/shuffle", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	response := ShuffleResponse{}
	if assert.NoError(t, shuffle(c)) {
		if assert.NoError(t, json.NewDecoder(rec.Body).Decode(&response)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, response.Results, "Sword", "Axe", "Bow")
		}
	}
}

func TestPick(t *testing.T) {
	e := echo.New()
	body := `{ "values": ["Sword", "Axe", "Bow"] }`
	req := httptest.NewRequest(http.MethodPost, "/pick", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	response := PickResponse{}
	if assert.NoError(t, pick(c)) {
		if assert.NoError(t, json.NewDecoder(rec.Body).Decode(&response)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Len(t, response.Results, 1)
		}
	}
}

func TestPickWithLimit(t *testing.T) {
	e := echo.New()
	body := `{ "values": ["Sword", "Axe", "Bow"], "limit": 2 }`
	req := httptest.NewRequest(http.MethodPost, "/pick", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	response := PickResponse{}
	if assert.NoError(t, pick(c)) {
		if assert.NoError(t, json.NewDecoder(rec.Body).Decode(&response)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Len(t, response.Results, 2)
		}
	}
}

func TestRoll(t *testing.T) {
	e := echo.New()
	body := `{ "expr": "d20 + 3" }`
	req := httptest.NewRequest(http.MethodPost, "/roll", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	response := RollResponse{}
	if assert.NoError(t, roll(c)) {
		if assert.NoError(t, json.NewDecoder(rec.Body).Decode(&response)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Greater(t, response.Result, 3.0)
		}
	}
}
