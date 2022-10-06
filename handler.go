package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type ShuffleRequest struct {
	Values []string `json:"values" query:"values" validate:"required"`
	Format string   `json:"format" query:"format"`
}

type ShuffleResponse struct {
	Result []string `json:"result"`
}

type PickRequest struct {
	Values []string `json:"values" query:"values" validate:"required"`
	Count  int      `json:"count" query:"count"`
	Format string   `json:"format" query:"format"`
}

type PickResponse struct {
	Result []string `json:"result"`
}

type RollRequest struct {
	Expression string `json:"expr" query:"expr" validate:"required"`
	Format     string `json:"format" query:"format"`
}

type RollResponse struct {
	Result float64 `json:"result"`
}

type MarkovRequest struct {
	Words     []string `json:"words" query:"words" validate:"required"`
	Order     int      `json:"order" query:"order"`
	Separator string   `json:"separator" query:"separator"`
	Count     int      `json:"count" query:"count"`
	Format    string   `json:"format" query:"format"`
}

type MarkovResponse struct {
	Result []string `json:"result"`
}

var badFormatFieldResponse = echo.NewHTTPError(http.StatusBadRequest, "Format field must be plain or json")

func shuffleHandler(c echo.Context) error {
	req := new(ShuffleRequest)
	req.Format = "json"
	if err := c.Bind(req); err != nil {
		return err
	}
	if !isCorrectFormat(req.Format) {
		return badFormatFieldResponse
	}
	if len(req.Values) <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Values field must contain at least one value")
	}

	res := shuffle(req.Values)

	if req.Format == "plain" {
		return c.String(http.StatusOK, strings.Join(res, ","))
	}
	return c.JSON(http.StatusOK, ShuffleResponse{Result: res})
}

func pickHandler(c echo.Context) error {
	req := new(PickRequest)
	req.Count = 1
	req.Format = "json"
	if err := c.Bind(req); err != nil {
		return err
	}
	if !isCorrectFormat(req.Format) {
		return badFormatFieldResponse
	}
	if req.Count <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Count field must be greater than zero")
	}
	if len(req.Values) <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Values field must contain at least one value")
	}

	results := pick(req.Values, req.Count)

	if req.Format == "plain" {
		return c.String(http.StatusOK, strings.Join(results, ","))
	}
	return c.JSON(http.StatusOK, PickResponse{Result: results})
}

func rollHandler(c echo.Context) error {
	req := new(RollRequest)
	req.Format = "json"
	if err := c.Bind(req); err != nil {
		return err
	}
	if !isCorrectFormat(req.Format) {
		return badFormatFieldResponse
	}

	res, err := roll(req.Expression)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect expression")
	}
	if req.Format == "plain" {
		return c.String(http.StatusOK, fmt.Sprintf("%.0f", res))
	}
	return c.JSON(http.StatusOK, RollResponse{Result: res})
}

func markovHandler(c echo.Context) error {
	req := new(MarkovRequest)
	req.Order = 1
	req.Count = 1
	req.Separator = ""
	req.Format = "json"
	if err := c.Bind(req); err != nil {
		return err
	}
	if !isCorrectFormat(req.Format) {
		return badFormatFieldResponse
	}
	if len(req.Words) <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Words field must contain at least one word")
	}
	if req.Count <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Count field must be greater than zero")
	}

	results := markov(req.Words, req.Order, req.Separator, req.Count)

	if req.Format == "plain" {
		return c.String(http.StatusOK, strings.Join(results, ","))
	}
	return c.JSON(http.StatusOK, MarkovResponse{Result: results})
}

func isCorrectFormat(format string) bool {
	return format == "plain" || format == "json"
}
