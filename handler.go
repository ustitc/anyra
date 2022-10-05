package main

import (
	"math/rand"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mb-14/gomarkov"
)

type ShuffleRequest struct {
	Values []string `json:"values" query:"values" validate:"required"`
}

type ShuffleResponse struct {
	Result []string `json:"result"`
}

type PickRequest struct {
	Values []string `json:"values" query:"values" validate:"required"`
	Count  int      `json:"count" query:"count"`
}

type PickResponse struct {
	Result []string `json:"result"`
}

type RollRequest struct {
	Expression string `json:"expr" query:"expr" validate:"required"`
}

type RollResponse struct {
	Result float64 `json:"result"`
}

type MarkovRequest struct {
	Words     []string `json:"words" query:"words" validate:"required"`
	Order     int      `json:"order" query:"order"`
	Separator string   `json:"separator" query:"separator"`
	Count     int      `json:"count" query:"count"`
}

type MarkovResponse struct {
	Result []string `json:"result"`
}

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
	return c.JSON(http.StatusOK, ShuffleResponse{Result: req.Values})
}

func pick(c echo.Context) error {
	req := new(PickRequest)
	req.Count = 1
	if err := c.Bind(req); err != nil {
		return err
	}
	if req.Count <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Count field must be greater than zero")
	}
	if len(req.Values) <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Values field must contain at least one value")
	}
	results := make([]string, req.Count)
	for i := 0; i < int(req.Count); i++ {
		results[i] = req.Values[rand.Intn(len(req.Values))]
	}
	return c.JSON(http.StatusOK, PickResponse{Result: results})
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

func markov(c echo.Context) error {
	req := new(MarkovRequest)
	req.Order = 1
	req.Count = 1
	req.Separator = ""
	if err := c.Bind(req); err != nil {
		return err
	}
	if len(req.Words) <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Words field must contain at least one word")
	}
	if req.Count <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Count field must be greater than zero")
	}

	chain := gomarkov.NewChain(req.Order)
	for _, word := range req.Words {
		chain.Add(strings.Split(word, req.Separator))
	}

	results := make([]string, req.Count)
	for i := 0; i < int(req.Count); i++ {
		results[i] = generate(chain, req.Separator)
	}
	return c.JSON(http.StatusOK, MarkovResponse{Result: results})
}

func generate(chain *gomarkov.Chain, sep string) string {
	order := chain.Order
	tokens := make([]string, 0)
	for i := 0; i < order; i++ {
		tokens = append(tokens, gomarkov.StartToken)
	}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := chain.Generate(tokens[(len(tokens) - order):])
		tokens = append(tokens, next)
	}
	return strings.Join(tokens[order:len(tokens)-1], sep)
}
