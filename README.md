# Anyra

**Anyra** (any random) - server for any kind of random generation you need.

To run a server:

``` bash
anyra run
```

This will start a server on default port `8080`.

To compile from source:

``` bash
go build
```

## `/pick`

Picks any value from a list.

**Example**

``` bash
curl 'http://localhost:8080/pick?values=Sword&values=Axe&values=Bow'
```

or

``` bash
curl \
  -X POST 'http://localhost:8080/pick' \
  -H 'Content-Type: application/json' \
  --data-binary '{ "values": ["Sword", "Axe", "Bow"] }'
```

Response: `200 OK`

``` json
{
    "result": [
        "Bow"
    ]
}
```


## `/shuffle`

Shuffles list of values.

**Example**

``` bash
curl 'http://localhost:8080/shuffle?values=Sword&values=Axe&values=Bow'
```

or

``` bash
curl \
  -X POST 'http://localhost:8080/shuffle' \
  -H 'Content-Type: application/json' \
  --data-binary '{ "values": ["Sword", "Axe", "Bow"] }'
```

Response: `200 OK`

``` json
{
    "result": [
        "Axe",
        "Bow",
        "Sword"
    ]
}
```

## `/roll`

Uses [dice syntax](https://en.wikipedia.org/wiki/Dice_notation) to generate random numbers.

**Example**

``` bash
curl 'http://localhost:8080/roll?expr=d20%2B3'
```

or

``` bash
curl \
  -X POST 'http://localhost:8080/roll' \
  -H 'Content-Type: application/json' \
  --data-binary '{ "expr": "d20 + 3" }'
```

Response: `200 OK`

``` json
{
    "result": 17
}
```

## `/markov`

TODO
