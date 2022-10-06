package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2"
)

var countFlag = &cli.IntFlag{
	Name:    "count",
	Aliases: []string{"c"},
	Value:   1,
	Usage:   "number of values to generate",
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "server",
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
			{
				Name:  "shuffle",
				Usage: "shuffle values",
				Action: func(cCtx *cli.Context) error {
					values := cCtx.Args().Slice()
					if len(values) == 0 {
						return cli.Exit(`"anyra shuffle" requires at least 1 argument`, 1)
					}
					res := shuffle(values)
					fmt.Println(strings.Join(res, ", "))
					return nil
				},
			},
			{
				Name:  "pick",
				Usage: "pick a value",
				Flags: []cli.Flag{
					countFlag,
				},
				Action: func(cCtx *cli.Context) error {
					values := cCtx.Args().Slice()
					if len(values) == 0 {
						return cli.Exit(`"anyra pick" requires at least 1 argument`, 1)
					}
					res := pick(values, cCtx.Int("count"))
					fmt.Println(strings.Join(res, ", "))
					return nil
				},
			},
			{
				Name:  "roll",
				Usage: "roll a dice",
				Action: func(cCtx *cli.Context) error {
					args := cCtx.Args().Slice()
					if len(args) == 0 {
						return cli.Exit(`"anyra roll" requires at least 1 argument`, 1)
					}
					expr := strings.Join(args, "")
					res, err := roll(expr)
					if err != nil {
						return cli.Exit(`Passed expression has wrong format`, 1)
					}
					fmt.Println(res)
					return nil
				},
			},
			{
				Name:  "markov",
				Usage: "create and evaluate markov chain",
				Flags: []cli.Flag{
					countFlag,
					&cli.IntFlag{
						Name:    "order",
						Aliases: []string{"o"},
						Value:   1,
						Usage:   "order of markov chain",
					},
					&cli.StringFlag{
						Name:    "separator",
						Aliases: []string{"s"},
						Value:   "",
						Usage:   "separator with which to divide words",
					},
				},
				Action: func(cCtx *cli.Context) error {
					words := cCtx.Args().Slice()
					if len(words) == 0 {
						return cli.Exit(`"anyra markov" requires at least 1 argument`, 1)
					}
					res := markov(words, cCtx.Int("order"), cCtx.String("separator"), cCtx.Int("count"))
					fmt.Println(strings.Join(res, ", "))
					return nil
				},
			},
		},
		Name:  "anyra",
		Usage: "tool for any kind of random generation",
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

	e.GET("/shuffle", shuffleHandler)
	e.POST("/shuffle", shuffleHandler)
	e.POST("/pick", pickHandler)
	e.GET("/pick", pickHandler)
	e.GET("/roll", rollHandler)
	e.POST("/roll", rollHandler)
	e.GET("/markov", markovHandler)
	e.POST("/markov", markovHandler)
	e.Logger.Fatal(e.Start(":" + port))
}
