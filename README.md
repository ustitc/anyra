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

**Request Parameters**

|Name|Type|Required|Description|
|----|----|--------|-----------|
|values|[]string|true|Values to pick from|
|count|int|false|Number of values to pick|

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

**Request Parameters**

|Name|Type|Required|Description|
|----|----|--------|-----------|
|values|[]string|true|Values to shuffle|

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

|Name|Type|Required|Description|
|----|----|--------|-----------|
|expr|string|true|Dice expression|

## `/markov`

``` bash
curl 'http://localhost:8080/markov?words=Gayel&words=Broggs&words=Deron&words=Jenzen&words=Adryan&words=Damaris&words=Ragan&words=Rodrock&words=Chindler'
```

or

``` bash
curl \
  -X POST 'http://localhost:8080/roll' \
  -H 'Content-Type: application/json' \
  --data-binary '{ "words": ["Gayel", "Broggs", "Deron", "Jenzen", "Adryan", "Damaris", "Ragan", "Rodrock", "Chindler"] }'
```

Response: `200 OK`

``` json
{
  "result":[
    "Rarodris"
  ]
}
```

|Name|Type|Required|Description|
|----|----|--------|-----------|
|words|[]string|true|Words for markov chain|
|count|int|false|Number of words to generate|
