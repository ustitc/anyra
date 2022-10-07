# Anyra

[![Go Report Card](https://goreportcard.com/badge/github.com/ustits/anyra)](https://goreportcard.com/report/github.com/ustits/anyra)

---

**Anyra** (any random) - tool for any kind of random generation you need.

## ✨ Features

- **Pick and shuffle** - choose random values or shuffle them all together
- **Dice notation** - roll `d20` for your ability checks
- **Markov chain** - generate new names or plot hooks with the help of Markov chains
- **[HTTP API](./docs/api.md)** - run anyra as a server and get all listed features via http

## 🚀 Quick start

Compile from source:

``` bash
go build
```

List available commands:

``` bash
anyra help
```

To run a server on default `8080` port:

``` bash
anyra server
```
