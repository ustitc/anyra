package main

import (
	"math/rand"
	"strings"

	"github.com/mb-14/gomarkov"
)

func shuffle(values []string) []string {
	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})
	return values
}

func pick(values []string, count int) []string {
	results := make([]string, count)
	for i := 0; i < int(count); i++ {
		results[i] = values[rand.Intn(len(values))]
	}
	return results
}

func roll(expr string) (float64, error) {
	compiled, err := parser.ParseString("", expr)
	if err != nil {
		return 0, err
	}
	return compiled.Eval(), nil
}

func markov(words []string, order int, separator string, count int) []string {
	chain := gomarkov.NewChain(order)
	for _, word := range words {
		chain.Add(strings.Split(word, separator))
	}

	results := make([]string, count)
	for i := 0; i < int(count); i++ {
		results[i] = generate(chain, separator)
	}
	return results
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
