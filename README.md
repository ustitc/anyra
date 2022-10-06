# Anyra

[![Go Report Card](https://goreportcard.com/badge/github.com/ustits/anyra)](https://goreportcard.com/report/github.com/ustits/anyra)

---

**Anyra** (any random) - tool for any kind of random generation you need.

To run a server:

``` bash
anyra server
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

|Name|Type|Required|Default|Description|
|----|----|--------|-------|-----------|
|values|[]string|true|-|Values to pick from|
|count|int|false|1|Number of values to pick|
|format|string|false|`json`|What format to use in response (`plain`,`json`) |

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

|Name|Type|Required|Default|Description|
|----|----|--------|-------|-----------|
|values|[]string|true|-|Values to shuffle|
|format|string|false|`json`|What format to use in response (`plain`,`json`) |

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

**Request Parameters**

|Name|Type|Required|Default|Description|
|----|----|--------|-------|-----------|
|expr|string|true|-|Dice expression|
|format|string|false|`json`|What format to use in response (`plain`,`json`) |

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

**Request Parameters**

|Name|Type|Required|Default|Description|
|----|----|--------|-------|-----------|
|words|[]string|true|-|Words for markov chain|
|order|int|false|1|Order of markov chain|
|separator|string|false|""|Separator with which to divide words|
|count|int|false|1|Number of words to generate|
|format|string|false|`json`|What format to use in response (`plain`,`json`) |
