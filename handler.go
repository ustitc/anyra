package main

import (
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
)

func shuffle(c echo.Context) error {
	req := new(ShuffleRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if len(req.Values) <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Values field must contain at least one value")
	}
	rand.Shuffle(len(req.Values), func(i, j int) {
		req.Values[i], req.Values[j] = req.Values[j], req.Values[i]
	})
	return c.JSON(http.StatusOK, ShuffleResponse{Results: req.Values})
}

func pick(c echo.Context) error {
	req := new(PickRequest)
	req.Limit = 1
	if err := c.Bind(req); err != nil {
		return err
	}
	if req.Limit <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Limit field must be greater than zero")
	}
	if len(req.Values) <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Values field must contain at least one value")
	}
	results := make([]string, req.Limit)
	for i := 0; i < int(req.Limit); i++ {
		results[i] = req.Values[rand.Intn(len(req.Values))]
	}
	return c.JSON(http.StatusOK, PickResponse{Results: results})
}

func roll(c echo.Context) error {
	req := new(RollRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	expr, err := parser.ParseString("", req.Expression)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect expression")
	}
	res := expr.Eval()
	return c.JSON(http.StatusOK, RollResponse{Result: res})
}
