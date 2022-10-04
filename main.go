package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"text/scanner"
	"unicode"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2"
)

type ShuffleRequest struct {
	Values []string `json:"values" query:"values" validate:"required"`
}

type ShuffleResponse struct {
	Results []string `json:"result" query:"result"`
}

type PickRequest struct {
	Values []string `json:"values" query:"values" validate:"required"`
	Limit  int      `json:"limit" query:"limit"`
}

type PickResponse struct {
	Results []string `json:"result" query:"result"`
}

type RollRequest struct {
	Expression string `json:"expr" query:"expr" validate:"required"`
}

type RollResponse struct {
	Result float64 `json:"result" query:"result"`
}

var (
	lex = lexer.NewTextScannerLexer(func(s *scanner.Scanner) {
		// to parse d20 without whitespaces
		s.IsIdentRune = func(ch rune, i int) bool {
			return unicode.IsDigit(ch) && i > 0
		}
	})
	parser = participle.MustBuild[Expression](
		participle.Lexer(lex),
	)
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "run server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   "8080",
					},
				},
				Action: func(cCtx *cli.Context) error {
					runServer(cCtx.String("port"))
					return nil
				},
			},
		},
		Name:  "anyra",
		Usage: "server for any kind of random generation",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

func runServer(port string) {
	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/shuffle", shuffle)
	e.POST("/shuffle", shuffle)
	e.POST("/pick", pick)
	e.GET("/pick", pick)
	e.GET("/roll", roll)
	e.POST("/roll", roll)
	e.Logger.Fatal(e.Start(":" + port))
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
