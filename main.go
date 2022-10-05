package main

import (
	"log"
	"os"
	"text/scanner"
	"unicode"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2"
)

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
	e.GET("/markov", markov)
	e.POST("/markov", markov)
	e.Logger.Fatal(e.Start(":" + port))
}
