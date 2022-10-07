package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	countFlag = &cli.IntFlag{
		Name:    "count",
		Aliases: []string{"c"},
		Value:   1,
		Usage:   "number of values to generate",
	}
	fileFlag = &cli.StringFlag{
		Name:    "file",
		Aliases: []string{"f"},
		Usage:   "file with values separated by newline",
	}
	fileFailExit = cli.Exit(`Can't read file"`, 1)
)

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
				Flags: []cli.Flag{
					fileFlag,
				},
				Action: func(cCtx *cli.Context) error {
					values, err := readValues(cCtx)
					if err != nil {
						return fileFailExit
					}
					if len(values) == 0 {
						return cli.Exit(`"anyra shuffle" requires at least 1 value`, 1)
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
					fileFlag,
				},
				Action: func(cCtx *cli.Context) error {
					values, err := readValues(cCtx)
					if err != nil {
						return fileFailExit
					}
					if len(values) == 0 {
						return cli.Exit(`"anyra pick" requires at least 1 value`, 1)
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
					fileFlag,
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
					words, err := readValues(cCtx)
					if err != nil {
						return fileFailExit
					}
					if len(words) == 0 {
						return cli.Exit(`"anyra markov" requires at least 1 value`, 1)
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

func readValues(cCtx *cli.Context) ([]string, error) {
	file := cCtx.String("file")
	if file == "" {
		return cCtx.Args().Slice(), nil
	}
	return readLines(file)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
